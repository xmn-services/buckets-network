package filedisks

import (
	"path/filepath"

	application_chains "github.com/xmn-services/buckets-network/application/chains"
	application_miner "github.com/xmn-services/buckets-network/application/miners"
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
	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
	transfer_mined_link "github.com/xmn-services/buckets-network/domain/transfers/links/mined"
	libs_file "github.com/xmn-services/buckets-network/libs/file"
)

const chunksDirName = "chunks"
const bucketFilesDirName = "files"
const bucketsDirName = "buckets"
const genesisDirName = "genesis"
const blocksDirName = "blocks"
const minedBlocksDirName = "mined_blocks"
const linksDirName = "links"
const minedLinksDirName = "mined_links"
const chainsDirName = "chains"

// NewChainApplication returns a new chain application
func NewChainApplication(
	basePath string,
	genesisFileNameWithExt string,
	chainFileName string,
	chainFileExt string,
) application_chains.Application {
	minerApp := NewMinerApplication(
		basePath,
		genesisFileNameWithExt,
	)

	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	genesisService := NewGenesisService(basePath, genesisFileNameWithExt)

	chainRepository := NewChainRepository(basePath, chainFileName, chainFileExt, genesisRepository)
	chainService := NewChainService(basePath, chainFileName, chainFileExt, genesisRepository, genesisService)

	return application_chains.NewApplication(minerApp, chainRepository, chainService)
}

// NewMinerApplication creates a new miner application
func NewMinerApplication(
	basePath string,
	genesisFileNameWithExt string,
) application_miner.Application {
	bucketRepository := NewBucketRepository(basePath)
	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	minedBlockRepository := NewMinedBlockRepository(basePath, genesisRepository)
	return application_miner.NewApplication(bucketRepository, minedBlockRepository, genesisRepository)
}

// NewChunkRepository creates a new chunks repository
func NewChunkRepository(basePath string) chunks.Repository {
	path := filepath.Join(basePath, chunksDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	trChunkRepository := transfer_chunk.NewRepository(fileRepository)
	return chunks.NewRepository(trChunkRepository)
}

// NewChunkService creates a new chunks service
func NewChunkService(basePath string) chunks.Service {
	path := filepath.Join(basePath, chunksDirName)
	fileService := libs_file.NewFileDiskService(path)

	chunkRepository := NewChunkRepository(basePath)
	trChunkService := transfer_chunk.NewService(fileService)
	return chunks.NewService(chunkRepository, trChunkService)
}

// NewBucketFileRepository creates a new bucketFile repository
func NewBucketFileRepository(basePath string) bucket_files.Repository {
	path := filepath.Join(basePath, bucketFilesDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	chunkRepository := NewChunkRepository(basePath)
	trBucketFileRepository := transfer_file.NewRepository(fileRepository)
	return bucket_files.NewRepository(chunkRepository, trBucketFileRepository)
}

// NewBucketFileService creates a new bucketFile service
func NewBucketFileService(basePath string) bucket_files.Service {
	path := filepath.Join(basePath, bucketFilesDirName)
	fileService := libs_file.NewFileDiskService(path)

	chunkService := NewChunkService(basePath)
	bucketFileRepository := NewBucketFileRepository(basePath)
	trBucketFileService := transfer_file.NewService(fileService)
	return bucket_files.NewService(chunkService, bucketFileRepository, trBucketFileService)
}

// NewBucketRepository creates a new bucket repository
func NewBucketRepository(basePath string) buckets.Repository {
	path := filepath.Join(basePath, bucketsDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	bucketFileRepository := NewBucketFileRepository(basePath)
	trBucketRepository := transfer_bucket.NewRepository(fileRepository)
	return buckets.NewRepository(bucketFileRepository, trBucketRepository)
}

// NewBucketService returns the bucket service
func NewBucketService(basePath string) buckets.Service {
	path := filepath.Join(basePath, bucketsDirName)
	fileService := libs_file.NewFileDiskService(path)

	bucketFileService := NewBucketFileService(basePath)
	bucketRepository := NewBucketRepository(basePath)
	trBucketService := transfer_bucket.NewService(fileService)
	return buckets.NewService(bucketFileService, bucketRepository, trBucketService)
}

// NewGenesisRepository returns a new genesis repository
func NewGenesisRepository(basePath string, genesisFileNameWithExt string) genesis.Repository {
	path := filepath.Join(basePath, genesisDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	trGenesisRepository := transfer_genesis.NewRepository(fileRepository, genesisFileNameWithExt)
	return genesis.NewRepository(trGenesisRepository)
}

// NewGenesisService creates a new genesis service
func NewGenesisService(basePath string, genesisFileNameWithExt string) genesis.Service {
	path := filepath.Join(basePath, genesisDirName)
	fileService := libs_file.NewFileDiskService(path)

	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	trGenesisService := transfer_genesis.NewService(fileService, genesisFileNameWithExt)
	return genesis.NewService(genesisRepository, trGenesisService)
}

// NewBlockRepository creates a new block repository
func NewBlockRepository(basePath string, genesisRepository genesis.Repository) blocks.Repository {
	path := filepath.Join(basePath, blocksDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	bucketRepository := NewBucketRepository(basePath)
	trBlockRepository := transfer_block.NewRepository(fileRepository)
	return blocks.NewRepository(genesisRepository, bucketRepository, trBlockRepository)
}

// NewBlockService creates a new block service
func NewBlockService(basePath string, genesisRepository genesis.Repository, genesisService genesis.Service) blocks.Service {
	path := filepath.Join(basePath, blocksDirName)
	fileService := libs_file.NewFileDiskService(path)

	blockRepository := NewBlockRepository(basePath, genesisRepository)
	bucketService := NewBucketService(basePath)
	trBlockService := transfer_block.NewService(fileService)
	return blocks.NewService(blockRepository, bucketService, trBlockService)
}

// NewMinedBlockRepository creates a new mined block repository
func NewMinedBlockRepository(basePath string, genesisRepository genesis.Repository) mined_block.Repository {
	path := filepath.Join(basePath, minedBlocksDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	blockRepository := NewBlockRepository(basePath, genesisRepository)
	trMinedBlockRepository := transfer_block_mined.NewRepository(fileRepository)
	return mined_block.NewRepository(blockRepository, trMinedBlockRepository)
}

// NewMinedBlockService creates a new mined block service
func NewMinedBlockService(basePath string, genesisRepository genesis.Repository, genesisService genesis.Service) mined_block.Service {
	path := filepath.Join(basePath, minedBlocksDirName)
	fileService := libs_file.NewFileDiskService(path)

	minedBlockRepository := NewMinedBlockRepository(basePath, genesisRepository)
	blockService := NewBlockService(basePath, genesisRepository, genesisService)
	trMinedBlockService := transfer_block_mined.NewService(fileService)
	return mined_block.NewService(minedBlockRepository, blockService, trMinedBlockService)
}

// NewLinkRepository creates a new link repository
func NewLinkRepository(basePath string, genesisRepository genesis.Repository) links.Repository {
	path := filepath.Join(basePath, linksDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	minedBlockRepository := NewMinedBlockRepository(basePath, genesisRepository)
	trLinkRepository := transfer_link.NewRepository(fileRepository)
	return links.NewRepository(minedBlockRepository, trLinkRepository)
}

// NewLinkService creates a new link service
func NewLinkService(basePath string, genesisRepository genesis.Repository, genesisService genesis.Service) links.Service {
	path := filepath.Join(basePath, linksDirName)
	fileService := libs_file.NewFileDiskService(path)

	linkRepository := NewLinkRepository(basePath, genesisRepository)
	minedBlockService := NewMinedBlockService(basePath, genesisRepository, genesisService)
	trLinkService := transfer_link.NewService(fileService)
	return links.NewService(linkRepository, minedBlockService, trLinkService)
}

// NewMinedLinkRepository creates a new mined link repository
func NewMinedLinkRepository(basePath string, genesisRepository genesis.Repository) mined_link.Repository {
	path := filepath.Join(basePath, minedLinksDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	linkRepository := NewLinkRepository(basePath, genesisRepository)
	trMinedLinkRepository := transfer_mined_link.NewRepository(fileRepository)
	return mined_link.NewRepository(linkRepository, trMinedLinkRepository)
}

// NewMinedLinkService creates a new mined link service
func NewMinedLinkService(basePath string, genesisRepository genesis.Repository, genesisService genesis.Service) mined_link.Service {
	path := filepath.Join(basePath, minedLinksDirName)
	fileService := libs_file.NewFileDiskService(path)

	minedLinkRepository := NewMinedLinkRepository(basePath, genesisRepository)
	linkService := NewLinkService(basePath, genesisRepository, genesisService)
	trMinedLinkService := transfer_mined_link.NewService(fileService)
	return mined_link.NewService(minedLinkRepository, linkService, trMinedLinkService)
}

// NewChainRepository creates a new chain repository
func NewChainRepository(basePath string, chainFileName string, chainFileExt string, genesisRepository genesis.Repository) chains.Repository {
	path := filepath.Join(basePath, chainsDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	minedBlockRepository := NewMinedBlockRepository(basePath, genesisRepository)
	minedLinkRepository := NewMinedLinkRepository(basePath, genesisRepository)
	trChainRepository := transfer_chains.NewRepository(fileRepository, chainFileName, chainFileExt)
	return chains.NewRepository(genesisRepository, minedBlockRepository, minedLinkRepository, trChainRepository)
}

// NewChainService creates a new chain service
func NewChainService(basePath string, chainFileName string, chainFileExt string, genesisRepository genesis.Repository, genesisService genesis.Service) chains.Service {
	path := filepath.Join(basePath, chainsDirName)
	fileService := libs_file.NewFileDiskService(path)

	chainRepository := NewChainRepository(basePath, chainFileName, chainFileExt, genesisRepository)
	minedBlockRepository := NewMinedBlockRepository(basePath, genesisRepository)
	minedBlockService := NewMinedBlockService(basePath, genesisRepository, genesisService)
	minedLinkRepository := NewMinedLinkRepository(basePath, genesisRepository)
	minedLinkService := NewMinedLinkService(basePath, genesisRepository, genesisService)
	trChainService := transfer_chains.NewService(fileService, chainFileName, chainFileExt)
	return chains.NewService(chainRepository, genesisService, minedBlockRepository, minedBlockService, minedLinkRepository, minedLinkService, trChainService)
}
