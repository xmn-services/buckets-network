package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	bucket_files "github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
	transfer_block_mined "github.com/xmn-services/buckets-network/domain/transfers/blocks/mined"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	transfer_file "github.com/xmn-services/buckets-network/domain/transfers/buckets/files"
	transfer_chunk "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks"
	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
	libs_file "github.com/xmn-services/buckets-network/libs/file"
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
	fileRepository libs_file.Repository,
	fileService libs_file.Service,
	genesisFileNameWithExt string,
) Application {
	hashAdapter := hash.NewAdapter()

	trChunkRepository := transfer_chunk.NewRepository(fileRepository)
	chunkRepository := chunks.NewRepository(trChunkRepository)

	trBucketFileRepository := transfer_file.NewRepository(fileRepository)
	bucketFileRepository := bucket_files.NewRepository(chunkRepository, trBucketFileRepository)

	trBucketRepository := transfer_bucket.NewRepository(fileRepository)
	bucketRepository := buckets.NewRepository(bucketFileRepository, trBucketRepository)

	genesisBuilder := genesis.NewBuilder()
	trGenesisRepository := transfer_genesis.NewRepository(fileRepository, genesisFileNameWithExt)
	genesisRepository := genesis.NewRepository(trGenesisRepository)

	blockBuilder := blocks.NewBuilder()
	trBlockRepository := transfer_block.NewRepository(fileRepository)
	blockRepository := blocks.NewRepository(genesisRepository, bucketRepository, trBlockRepository)

	minedBlockBuilder := mined_block.NewBuilder()
	trMinedBlockRepository := transfer_block_mined.NewRepository(fileRepository)
	minedBlockRepository := mined_block.NewRepository(blockRepository, trMinedBlockRepository)

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
