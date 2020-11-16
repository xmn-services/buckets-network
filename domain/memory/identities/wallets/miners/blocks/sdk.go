package blocks

import (
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
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
	return createBuilder()
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
	WithBlocks(blocks []mined_blocks.Block) Builder
	Now() (Blocks, error)
}

// Blocks represents a blocks
type Blocks interface {
	All() []mined_blocks.Block
	Add(block mined_blocks.Block) error
	Delete(hash hash.Hash) error
}
