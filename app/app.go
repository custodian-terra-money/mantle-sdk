package app

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"
	TerraApp "github.com/terra-project/core/app"
	compatapp "github.com/terra-project/mantle-compatibility/app"
	"github.com/terra-project/mantle-sdk/app/mantlemint"
	"github.com/terra-project/mantle-sdk/app/middlewares"
	"github.com/terra-project/mantle-sdk/db/force_transaction"
	"github.com/terra-project/mantle-sdk/utils"
	"log"
	"reflect"
	"time"

	"github.com/terra-project/mantle-sdk/committer"

	"github.com/terra-project/mantle-sdk/db"
	"github.com/terra-project/mantle-sdk/depsresolver"
	"github.com/terra-project/mantle-sdk/graph"
	"github.com/terra-project/mantle-sdk/graph/schemabuilders"
	"github.com/terra-project/mantle-sdk/indexer"
	"github.com/terra-project/mantle-sdk/querier"
	reg "github.com/terra-project/mantle-sdk/registry"
	"github.com/terra-project/mantle-sdk/subscriber"
	"github.com/terra-project/mantle-sdk/types"
)

type Mantle struct {
	isSynced             bool
	app                  *TerraApp.TerraApp
	registry             *reg.Registry
	mantlemint           mantlemint.Mantlemint
	gqlInstance          *graph.GraphQLInstance
	depsResolverInstance depsresolver.DepsResolver
	committerInstance    committer.Committer
	indexerInstance      *indexer.IndexerBaseInstance
	db                   db.DBwithGlobalTransaction
}

type SyncConfiguration struct {
	TendermintEndpoint string
	SyncUntil          uint64
	Reconnect          bool
	OnWSError          func(err error)
	OnInjectError      func(err error)
}

var (
	GlobalTerraApp *TerraApp.TerraApp
)

func NewMantle(
	db db.DB,
	genesis *tmtypes.GenesisDoc,
	indexers ...types.IndexerRegisterer,
) (mantleApp *Mantle) {
	// wrap db w/ force transaction manager
	dbWithTransaction := force_transaction.WithGlobalTransactionManager(db)

	// create new terra app with postgres-patched KVStore
	tmdb := dbWithTransaction.GetCosmosAdapter()
	terraApp := compatapp.NewTerraApp(tmdb)
	GlobalTerraApp = terraApp

	// gather outputs of indexer registry
	registry := reg.NewRegistry(indexers)

	// initialize gql
	depsResolverInstance := depsresolver.NewDepsResolver()
	querierInstance := querier.NewQuerier(db, registry.KVIndexMap)

	// instantiate gql
	gqlInstance := graph.NewGraphQLInstance(
		depsResolverInstance,
		querierInstance,
		schemabuilders.CreateABCIStubSchemaBuilder(terraApp),
		schemabuilders.CreateMantleStateSchemaBuilder(nil, nil),
		schemabuilders.CreateModelSchemaBuilder(nil, reflect.TypeOf((*types.BlockState)(nil))),
		schemabuilders.CreateModelSchemaBuilder(registry.KVIndexMap, registry.Models...),
	)

	// initialize committer
	committerInstance := committer.NewCommitter(db, registry.KVIndexMap)

	// initialize indexer
	indexerInstance := indexer.NewIndexerBaseInstance(
		registry.Indexers,
		registry.IndexerOutputs,
		gqlInstance.QueryInternal,
		gqlInstance.Commit,
	)

	mantleApp = &Mantle{
		isSynced: false,
		app:      terraApp,
		registry: nil,
		mantlemint: mantlemint.NewMantlemint(
			tmdb,
			terraApp,

			// in order to prevent indexer output to be in disparity
			// w/ tendermint state, indexers MUST run before commit.
			// use middlewares to run indexers (in CommitSync/CommitAsync)
			middlewares.NewIndexerMiddleware(func(responses state.ABCIResponses) {
				mantleApp.indexerLifecycle(responses)
			}),
		),
		gqlInstance:          gqlInstance,
		depsResolverInstance: depsResolverInstance,
		committerInstance:    committerInstance,
		indexerInstance:      indexerInstance,
		db:                   dbWithTransaction,
	}

	// initialize chain
	if initErr := mantleApp.mantlemint.Init(genesis); initErr != nil {
		panic(initErr)
	}

	return mantleApp
}

func (mantle *Mantle) indexerLifecycle(responses state.ABCIResponses) {
	var block = mantle.mantlemint.GetCurrentBlock()
	var height = block.Header.Height

	// create blockState
	var deliverTxsCopy = make([]abci.ResponseDeliverTx, len(responses.DeliverTxs))
	for i, deliverTx := range responses.DeliverTxs {
		deliverTxsCopy[i] = *deliverTx
	}
	blockState := types.BlockState{
		Height:             block.Height,
		ResponseBeginBlock: *responses.BeginBlock,
		ResponseEndBlock:   *responses.EndBlock,
		ResponseDeliverTx:  deliverTxsCopy,
		Block:              utils.ConvertBlockToRawBlock(block),
	}

	// set BlockState in depsResolver
	mantle.depsResolverInstance.SetPredefinedState(blockState)

	// RunIndexerRound panics when an indexer fails
	mantle.indexerInstance.RunIndexerRound()

	// flush states to database
	// note that indexer outputs are committed __BEFORE__ IAVL
	// because reversing indexer outputs is trivial (i.e. overwrite them)
	// whereas IAVL reversal is a little tricky.
	indexerOutputs := mantle.depsResolverInstance.GetState()
	defer mantle.depsResolverInstance.Dispose()

	// convert indexer outputs to slice
	var commitTargets = make([]interface{}, len(indexerOutputs))
	var i = 0
	for _, output := range indexerOutputs {
		commitTargets[i] = output
		i++
	}

	if commitErr := mantle.committerInstance.Commit(uint64(height), commitTargets...); commitErr != nil {
		panic(commitErr)
	}
}

func (mantle *Mantle) QuerySync(configuration SyncConfiguration, currentBlockHeight int64) {
	log.Println("Local blockchain is behind, syncing previous blocks...")
	remoteBlock, err := subscriber.GetBlock(fmt.Sprintf("http://%s/block", configuration.TendermintEndpoint))

	if err != nil {
		panic(fmt.Errorf("error during mantle sync: remote head fetch failed. fromHeight=%d, (%s)", currentBlockHeight, err))
	}

	remoteHeight := remoteBlock.Header.Height
	syncingBlockHeight := currentBlockHeight
	tStart := time.Now()

	for syncingBlockHeight < remoteHeight {
		// stop sync if SyncUntil is given
		if configuration.SyncUntil != 0 && uint64(syncingBlockHeight) == configuration.SyncUntil {
			for {
			}
		}

		remoteBlock, err := subscriber.GetBlock(fmt.Sprintf("http://%s/block?height=%d", configuration.TendermintEndpoint, syncingBlockHeight+1))
		if err != nil {
			panic(fmt.Errorf("error during mantle sync: remote block(%d) fetch failed", syncingBlockHeight))
		}

		// run round
		if _, err := mantle.Inject(remoteBlock); err != nil {
			if configuration.OnInjectError != nil {
				configuration.OnInjectError(err)
			} else {
				panic(err)
			}
		}

		syncingBlockHeight++
	}

	dur := time.Now().Sub(tStart)

	if dur > time.Second {
		log.Printf("[mantle] QuerySync: %d to %d, Elapsed: %dms", currentBlockHeight, remoteHeight, dur.Milliseconds())
	}
}

func (mantle *Mantle) Sync(configuration SyncConfiguration) {
	// subscribe to NewBlock event
	rpcSubscription, connRefused := subscriber.NewRpcSubscription(
		fmt.Sprintf("ws://%s/websocket", configuration.TendermintEndpoint),
		configuration.OnWSError,
	)

	// connRefused here is most likely triggered by ECONNREFUSED
	// in case reconnect flag is set, try reestablish the connection after 5 seconds.
	if connRefused != nil {
		if configuration.Reconnect {
			select {
				case <-time.NewTimer(5 * time.Second).C:
					mantle.Sync(configuration)
			}
			return

		} else {
			panic(connRefused)
		}
	}

	blockChannel := rpcSubscription.Subscribe(configuration.Reconnect)

	for {
		select {
		case block := <-blockChannel:
			lastBlockHeight := mantle.app.LastBlockHeight()

			// stop sync if SyncUntil is given
			if configuration.SyncUntil != 0 && uint64(lastBlockHeight) == configuration.SyncUntil {
				for {
				}
			}

			if block.Header.Height-lastBlockHeight != 1 {
				log.Printf("lastBlockHeight=%v, remoteBlockHeight=%v\n", lastBlockHeight, block.Header.Height)
				mantle.QuerySync(configuration, lastBlockHeight)
			} else {
				if _, err := mantle.Inject(&block); err != nil {
					// if OnInjectError is set,
					// relay injection error to the caller
					if configuration.OnInjectError != nil {
						configuration.OnInjectError(err)
					} else {
						panic(err)
					}
				}
			}
		}
	}
}

func (mantle *Mantle) Server(port int) {
	go mantle.gqlInstance.ServeHTTP(port)
}

func (mantle *Mantle) Inject(block *types.Block) (*types.BlockState, error) {
	if block.Header.Height % 1000 == 0 {
		log.Printf("[mantle] db compaction started")
		if compactionErr := mantle.db.Compact(); compactionErr != nil {
			panic(compactionErr)
		}
		log.Printf("[mantle] db compaction done")
	}

	tStart := time.Now()

	// set global transaction boundary for
	// tendermint, cosmos, mantle
	mantle.db.SetGlobalTransactionBoundary()

	// inject this block
	blockState, injectErr := mantle.mantlemint.Inject(block)

	// if injection was successful,
	// flush all to disk
	if injectErr != nil {
		return blockState, injectErr
	}

	// flush everything to db
	if flushErr := mantle.db.FlushGlobalTransactionBoundary(); flushErr != nil {
		panic(fmt.Errorf("mantle could not commit to db, error=%v", flushErr))
	}

	tEnd := time.Now()

	log.Printf(
		"[mantle] Indexing finished for block(%d), processed in %dms",
		block.Header.Height,
		tEnd.Sub(tStart).Milliseconds(),
	)

	return blockState, injectErr
}

func (mantle *Mantle) ExportStates() map[string]interface{} {
	return mantle.depsResolverInstance.GetState()
}
