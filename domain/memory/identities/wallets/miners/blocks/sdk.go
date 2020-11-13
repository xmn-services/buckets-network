package blocks

import (
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter returns a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	mutableBuilder := entities.NewMutableBuilder()
	return createBuilder(hashAdapter, mutableBuilder)
}

// Adapter represents the blocks adapter
type Adapter interface {
	ToJSON(ins Blocks) *JSONBlocks
	ToBlocks(js *JSONBlocks) (Blocks, error)
}

// Factory represents a blocks factory
type Factory interface {
	Create() (Blocks, error)
}

// Builder represents a blocks builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithBlocks(blocks []mined_blocks.Block) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Blocks, error)
}

// Blocks represents a blocks
type Blocks interface {
	entities.Mutable
	All() []mined_blocks.Block
	Add(block mined_blocks.Block) error
	Delete(hash hash.Hash) error
}
