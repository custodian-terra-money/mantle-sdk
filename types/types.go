package types

import (
	tmtypes "github.com/tendermint/tendermint/types"
)

type (
	Tx = tmtypes.Tx
	//Block      = tmtypes.Block
	Block = struct {
		Header     tmtypes.Header       `json:"header"`
		Data       tmtypes.Data         `json:"data"`
		Evidence   tmtypes.EvidenceData `json:"evidence"`
		LastCommit *tmtypes.Commit      `json:"last_commit"`
	}
	Header     = tmtypes.Header
	GenesisDoc = tmtypes.GenesisDoc
)

//
//type Header struct {
//	Version struct {
//		Block uint64 `json:"Block,string"`
//		App   uint64 `json:"App,string"`
//	} `json:"version"`
//	ChainID string `json:"chain_id"`
//	Height  int64  `json:"height,string"`
//	Time    string `json:"time"`
//	// NumTxs      int64  `json:"num_txs,string"` // removed as of tendermint 0.33
//	// TotalTxs    int64  `json:"total_txs,string"` // remove as of tendermint 0.33
//	LastBlockID struct {
//		Hash  string `json:"hash"`
//		Parts struct {
//			Total string `json:"total"`
//			Hash  string `json:"hash"`
//		} `json:"parts"`
//	} `json:"last_block_id"`
//	LastCommitHash     string `json:"last_commit_hash"`
//	DataHash           string `json:"data_hash"`
//	ValidatorsHash     string `json:"validators_hash"`
//	NextValidatorsHash string `json:"next_validators_hash"`
//	ConsensusHash      string `json:"consensus_hash"`
//	AppHash            string `json:"app_hash"`
//	LastResultsHash    string `json:"last_results_hash"`
//	EvidenceHash       string `json:"evidence_hash"`
//	ProposerAddress    string `json:"proposer_address"`
//}
//
//type ResultBeginBlock struct {
//	Events []struct {
//		Type       string `json:"type"`
//		Attributes []struct {
//			Key   string `json:"key"`
//			Value string `json:"value"`
//		} `json:"attributes"`
//	} `json:"events"`
//}

type ResultEndBlock struct {
	ValidatorUpdates []interface{} `json:"validator_updates"`
	Events           []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"events"`
}

type Events struct {
	RewardsAmount                  []string `json:"rewards.amount"`
	RewardsValidator               []string `json:"rewards.validator"`
	ExchangeRateUpdateExchangeRate []string `json:"exchange_rate_update.exchange_rate"`
	TransferRecipient              []string `json:"transfer.recipient"`
	TransferAmount                 []string `json:"transfer.amount"`
	MessageSender                  []string `json:"message.sender"`
	ProposerRewardValidator        []string `json:"proposer_reward.validator"`
	CommissionValidator            []string `json:"commission.validator"`
	ProposerRewardAmount           []string `json:"proposer_reward.amount"`
	CommissionAmount               []string `json:"commission.amount"`
	ExchangeRateUpdateDenom        []string `json:"exchange_rate_update.denom"`
	TmEvent                        []string `json:"tm.event"`
}
