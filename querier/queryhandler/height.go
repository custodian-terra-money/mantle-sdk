package queryhandler

import (
	"fmt"
	"github.com/terra-project/mantle/db/kvindex"
	"strings"

	"github.com/terra-project/mantle/db"
	"github.com/terra-project/mantle/utils"
)

type HeightResolver struct {
	db          db.DB
	entityName  string
	indexName   string
	prefixStart []byte
	prefixEnd   []byte
}

// seek resolver
func NewHeightResolver(
	db db.DB,
	_ *kvindex.KVIndex,
	entityName,
	indexName string,
	indexOption interface{},
) (QueryHandler, error) {
	if !strings.HasPrefix(indexName, "Height") {
		return nil, nil
	}

	switch indexOption.(type) {
	case []interface{}:
		heightRange, _ := indexOption.([]interface{})
		return &HeightResolver{
			db:          db,
			entityName:  entityName,
			indexName:   "Height",
			prefixStart: utils.LeToBe(uint64(heightRange[0].(int))),
			prefixEnd:   utils.LeToBe(uint64(heightRange[1].(int))),
		}, nil
	case int:
		height, _ := indexOption.(int)
		return &HeightResolver{
			db:          db,
			entityName:  entityName,
			indexName:   "Height",
			prefixStart: utils.LeToBe(uint64(height)),
			prefixEnd:   nil,
		}, nil
	default:
		return nil, fmt.Errorf("invalid height parameters, entityName=%s, indexOption=%v", entityName, indexOption)
	}
}

func (resolver HeightResolver) Resolve() (QueryHandlerIterator, error) {
	entityNameInBytes := []byte(resolver.entityName)
	indexNameInBytes := []byte(resolver.indexName)

	prefixGroup := utils.BuildIndexGroupPrefix(
		entityNameInBytes,
		indexNameInBytes,
	)

	prefixStart := utils.BuildIndexIteratorPrefix(
		entityNameInBytes,
		indexNameInBytes,
		resolver.prefixStart,
	)

	var prefixEnd []byte = nil
	if resolver.prefixEnd != nil {
		prefixEnd = utils.BuildIndexIteratorPrefix(
			entityNameInBytes,
			indexNameInBytes,
			resolver.prefixEnd,
		)
	}

	return NewHeightResolverIterator(
		entityNameInBytes,
		prefixGroup,
		prefixStart,
		prefixEnd,
		resolver.db.IndexIterator(
			prefixEnd,
			true,
		),
	), nil
}

type HeightResolverIterator struct {
	entityName  []byte
	prefixGroup []byte
	prefixStart []byte
	prefixEnd   []byte
	it          db.Iterator
}

func NewHeightResolverIterator(entityName, prefixGroup, prefixStart, prefixEnd []byte, it db.Iterator) QueryHandlerIterator {
	return &HeightResolverIterator{
		entityName:  entityName,
		prefixGroup: prefixGroup,
		prefixStart: prefixStart,
		prefixEnd:   prefixEnd,
		it:          it,
	}
}

func (resolver *HeightResolverIterator) Valid() bool {
	return resolver.it.Valid(resolver.prefixStart) ||
		resolver.it.Valid(resolver.prefixEnd)
}
func (resolver *HeightResolverIterator) Next() {
	resolver.it.Next()
}
func (resolver *HeightResolverIterator) Key() []byte {
	return utils.BuildDocumentKey(
		resolver.entityName,
		resolver.it.DocumentKey(),
	)
}
func (resolver *HeightResolverIterator) Close() {
	resolver.it.Close()
}
