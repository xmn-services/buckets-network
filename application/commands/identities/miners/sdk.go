package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// maxMiningValue represents the max mining value before adding another miner number to the slice
const maxMiningValue = 2147483647

// maxMiningTries represents the max mining characters to try before abandonning
const maxMiningTries = 2147483647

// maxDifficulty represents the max difficulty a block can have
const maxDifficulty = 127

// defaultMiningValue represents the default mining value
const defaultMiningValue = uint8(0)

// NewApplication creates a new application
func NewApplication(
	blockRepository blocks.Repository,
	linkRepository links.Repository,
) Application {
	hashAdapter := hash.NewAdapter()
	return createApplication(
		hashAdapter,
		blockRepository,
		linkRepository,
	)
}

// Application represents a miner application
type Application interface {
	Test(difficulty uint) (string, error)
	Block(blockHashStr string) (string, error)
	Link(linkHashStr string) (string, error)
}
