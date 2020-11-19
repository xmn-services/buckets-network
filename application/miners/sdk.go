package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// maxMiningValue represents the max mining value before adding another miner number to the slice
const maxMiningValue = 2147483647

// maxMiningTries represents the max mining characters to try before abandonning
const maxMiningTries = 2147483647

// maxDifficulty represents the max difficulty a block can have
const maxDifficulty = 127

// NewApplication creates a new application instance
func NewApplication(
	bucketRepository buckets.Repository,
	minedBlockRepository mined_block.Repository,
	genesisRepository genesis.Repository,
) Application {
	hashAdapter := hash.NewAdapter()
	genesisBuilder := genesis.NewBuilder()
	blockBuilder := blocks.NewBuilder()
	minedBlockBuilder := mined_block.NewBuilder()
	linkBuilder := links.NewBuilder()
	minedLinkBuilder := mined_link.NewBuilder()
	chainBuilder := chains.NewBuilder()
	return createApplication(
		hashAdapter,
		bucketRepository,
		blockBuilder,
		minedBlockBuilder,
		minedBlockRepository,
		linkBuilder,
		minedLinkBuilder,
		genesisBuilder,
		genesisRepository,
		chainBuilder,
	)
}

// Application represents a miner application
type Application interface {
	Init(
		miningValue uint8,
		baseDifficulty uint,
		increasePerBucket float64,
		linkDifficulty uint,
		rootAdditionalBuckets uint,
		headAdditionalBuckets uint,
	) (chains.Chain, error)

	Block(
		bucketHashes []string,
		baseDifficulty uint,
		increasePerBucket float64,
		additionalBuckets uint,
	) (mined_block.Block, error)

	Link(
		prevMinedBlockHasStr string,
		newMinedBlockHashStr string,
		linkDifficulty uint,
	) (mined_link.Link, error)
}
