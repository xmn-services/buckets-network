package blocks

import (
	"hash"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Factory represents a blocks factory
type Factory interface {
	Create() Blocks
}

// Blocks represents a blocks
type Blocks interface {
	entities.Mutable
	All() []mined_blocks.Block
	Add(block mined_blocks.Block) error
	Delete(hash hash.Hash) error
}
