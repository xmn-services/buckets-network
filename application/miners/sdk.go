package miners

import (
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

// maxMiningValue represents the max mining value before adding another miner number to the slice
const maxMiningValue = 2147483647

// maxMiningTries represents the max mining characters to try before abandonning
const maxMiningTries = 2147483647

// miningBeginValue represents the first value of the hash that is expected on mining
const miningBeginValue = 0

// Application represents a miner application
type Application interface {
	Init(baseDifficulty uint, increasePerBucket float64, linkDifficulty uint) (chains.Chain, error)
	Block(bucketHashes []string, baseDifficulty uint, increasePerBucket float64) (mined_block.Block, error)
	Link(prevMinedBlockHasStr string, newMinedBlockHashStr string, linkDifficulty uint) (mined_link.Link, error)
}
