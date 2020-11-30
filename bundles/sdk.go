package bundles

import (
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/xmn-services/buckets-network/application/commands"
	application_chains "github.com/xmn-services/buckets-network/application/commands/chains"
	application_identity "github.com/xmn-services/buckets-network/application/commands/identities"
	application_identity_buckets "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	application_identity_chains "github.com/xmn-services/buckets-network/application/commands/identities/chains"
	application_identity_miners "github.com/xmn-services/buckets-network/application/commands/identities/miners"
	application_identity_storages "github.com/xmn-services/buckets-network/application/commands/identities/storages"
	application_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/application/commands/storages"
	"github.com/xmn-services/buckets-network/application/servers"
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	bucket_files "github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
	transfer_block_mined "github.com/xmn-services/buckets-network/domain/transfers/blocks/mined"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	transfer_file "github.com/xmn-services/buckets-network/domain/transfers/buckets/files"
	transfer_chunk "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks"
	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
	transfer_content "github.com/xmn-services/buckets-network/domain/transfers/contents"
	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
	transfer_mined_link "github.com/xmn-services/buckets-network/domain/transfers/links/mined"
	restapis_server "github.com/xmn-services/buckets-network/infrastructure/restapis/servers"
	libs_file "github.com/xmn-services/buckets-network/libs/file"

	"github.com/xmn-services/buckets-network/infrastructure/restapis/clients"
	client_chains "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/chains"
	client_identities "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/identities"
	client_identities_buckets "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/identities/buckets"
	client_identities_chains "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/identities/chains"
	client_identities_miners "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/identities/miners"
	client_identities_storages "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/identities/storages"
	client_peers "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/peers"
	client_storages "github.com/xmn-services/buckets-network/infrastructure/restapis/clients/storages"
)

const chunksDirName = "chunks"
const bucketFilesDirName = "bucket_files"
const bucketsDirName = "buckets"
const genesisDirName = "genesis"
const blocksDirName = "blocks"
const minedBlocksDirName = "mined_blocks"
const linksDirName = "links"
const minedLinksDirName = "mined_links"
const chainsDirName = "chains"
const peersDirName = "peers"
const contentsDirName = "contents"

// NewRestAPIClient creates a new rest api client
func NewRestAPIClient(peer peer.Peer) commands.Application {
	commandIdentityBuilder := NewRESTAPIClientIdentityBuilder(peer)
	peerApp := client_peers.NewApplication(peer)
	chainApp := client_chains.NewApplication(peer)
	storageApp := client_storages.NewApplication(peer)
	return clients.NewApplication(
		commandIdentityBuilder,
		peerApp,
		chainApp,
		storageApp,
		peer,
	)
}

// NewRESTAPIClientIdentityBuilder creates a new rest api client identity builder
func NewRESTAPIClientIdentityBuilder(peer peer.Peer) application_identity.Builder {
	bucketBuilder := client_identities_buckets.NewBuilder(peer)
	chainBuilder := client_identities_chains.NewBuilder(peer)
	minerBuilder := client_identities_miners.NewBuilder(peer)
	storageBuilder := client_identities_storages.NewBuilder(peer)
	return client_identities.NewBuilder(
		bucketBuilder,
		storageBuilder,
		chainBuilder,
		minerBuilder,
		peer,
	)
}

// NewRestAPIServer creates a new rest api server
func NewRestAPIServer(
	cmdApp commands.Application,
	maxUploadFileSize int64,
	waitPeriod time.Duration,
	port uint,
) servers.Application {
	router := mux.NewRouter()
	return restapis_server.NewApplication(
		cmdApp,
		router,
		maxUploadFileSize,
		waitPeriod,
		port,
	)
}

// NewCommandApplication creates a new command application
func NewCommandApplication(
	basePath string,
	peerFileNameWithExt string,
	genesisFileNameWithExt string,
	chainFileName string,
	chainFileExt string,
	identityExt string,
	chunkSizeInBytes uint,
	encPKBitrate int,
) commands.Application {
	peerApp := NewPeerApplication(basePath, peerFileNameWithExt)
	chainApp := NewChainApplication(basePath, genesisFileNameWithExt, chainFileName, chainFileExt)
	storageApp := NewStorageApplication(basePath)
	identityAppBuilder := NewIdentityApplicationBuilder(basePath, identityExt, chainFileName, chainFileExt, genesisFileNameWithExt, chunkSizeInBytes, encPKBitrate)
	identityRepository := identities.NewRepository(basePath, identityExt)
	identityService := identities.NewService(basePath, identityExt)
	return commands.NewApplication(
		peerApp,
		chainApp,
		storageApp,
		identityAppBuilder,
		identityRepository,
		identityService,
	)
}

// NewIdentityApplicationBuilder creates a new identity application builder
func NewIdentityApplicationBuilder(
	basePath string,
	extension string,
	chainFileName string,
	chainFileExt string,
	genesisFileNameWithExt string,
	chunkSizeInBytes uint,
	encPKBitrate int,
) application_identity.Builder {
	minerApp := NewIdentityMinerApplication(basePath, genesisFileNameWithExt)
	bucketAppBuilder := NewIdentityBucketApplicationBuilder(basePath, extension, chunkSizeInBytes, encPKBitrate)
	storageAppBuilder := NewIdentityStorageApplicationBuilder(basePath, extension)
	chainBuilder := NewIdentityChainApplicationBuilder(basePath, extension, chainFileName, chainFileExt, genesisFileNameWithExt)
	identityRepository := identities.NewRepository(basePath, extension)
	identityService := identities.NewService(basePath, extension)
	return application_identity.NewBuilder(
		minerApp,
		bucketAppBuilder,
		storageAppBuilder,
		chainBuilder,
		identityRepository,
		identityService,
	)
}

// NewIdentityChainApplicationBuilder represents a new identity chain application builder
func NewIdentityChainApplicationBuilder(
	basePath string,
	extension string,
	chainFileName string,
	chainFileExt string,
	genesisFileNameWithExt string,
) application_identity_chains.Builder {
	minerApplication := NewIdentityMinerApplication(basePath, genesisFileNameWithExt)
	identityRepository := identities.NewRepository(basePath, extension)
	identityService := identities.NewService(basePath, extension)
	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	genesisService := NewGenesisService(basePath, genesisFileNameWithExt)
	blockService := NewBlockService(basePath, genesisRepository, genesisService)
	linkService := NewLinkService(basePath, genesisRepository, genesisService)
	chainRepository := NewChainRepository(basePath, chainFileName, chainFileExt, genesisRepository)
	chainService := NewChainService(basePath, chainFileName, chainFileExt, genesisRepository, genesisService)
	return application_identity_chains.NewBuilder(
		minerApplication,
		identityRepository,
		identityService,
		genesisRepository,
		genesisService,
		blockService,
		linkService,
		chainRepository,
		chainService,
	)
}

// NewIdentityMinerApplication creates a new identity miner application
func NewIdentityMinerApplication(
	basePath string,
	genesisFileNameWithExt string,
) application_identity_miners.Application {
	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	blockRepository := NewBlockRepository(basePath, genesisRepository)
	linkRepository := NewLinkRepository(basePath, genesisRepository)
	return application_identity_miners.NewApplication(
		blockRepository,
		linkRepository,
	)
}

// NewIdentityStorageApplicationBuilder creates a new identity storage application builder
func NewIdentityStorageApplicationBuilder(
	basePath string,
	extension string,
) application_identity_storages.Builder {
	identityRepository := identities.NewRepository(basePath, extension)
	bucketRepository := NewBucketRepository(basePath)
	contentService := NewContentService(basePath)
	return application_identity_storages.NewBuilder(identityRepository, bucketRepository, contentService)
}

// NewIdentityBucketApplicationBuilder creates a new identity bucket application builder
func NewIdentityBucketApplicationBuilder(
	basePath string,
	extension string,
	chunkSizeInBytes uint,
	encPKBitrate int,
) application_identity_buckets.Builder {
	bucketRepository := NewBucketRepository(basePath)
	identityRepository := identities.NewRepository(basePath, extension)
	identityService := identities.NewService(basePath, extension)
	return application_identity_buckets.NewBuilder(
		bucketRepository,
		identityRepository,
		identityService,
		chunkSizeInBytes,
		encPKBitrate,
	)

}

// NewStorageApplication represents the storage application
func NewStorageApplication(
	basePath string,
) storages.Application {
	bucketRepository := NewBucketRepository(basePath)
	contentRepository := NewContentRepository(basePath)
	return storages.NewApplication(bucketRepository, contentRepository)
}

// NewPeerApplication creates a new peer application
func NewPeerApplication(
	basePath string,
	fileNameWithExt string,
) application_peers.Application {
	repository := NewPeerRepository(basePath, fileNameWithExt)
	service := NewPeerService(basePath, fileNameWithExt)
	return application_peers.NewApplication(repository, service)
}

// NewChainApplication returns a new chain application
func NewChainApplication(
	basePath string,
	genesisFileNameWithExt string,
	chainFileName string,
	chainFileExt string,
) application_chains.Application {
	genesisRepository := NewGenesisRepository(basePath, genesisFileNameWithExt)
	chainRepository := NewChainRepository(basePath, chainFileName, chainFileExt, genesisRepository)
	return application_chains.NewApplication(chainRepository)
}

// NewPeerRepository creates a new peer repository
func NewPeerRepository(
	basePath string,
	fileNameWithExt string,
) peers.Repository {
	path := filepath.Join(basePath, peersDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	return peers.NewRepository(fileRepository, fileNameWithExt)
}

// NewPeerService creates a new peer service
func NewPeerService(
	basePath string,
	fileNameWithExt string,
) peers.Service {
	path := filepath.Join(basePath, peersDirName)
	fileService := libs_file.NewFileDiskService(path)

	return peers.NewService(fileService, fileNameWithExt)
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

// NewContentRepository creates a new content repository
func NewContentRepository(basePath string) contents.Repository {
	path := filepath.Join(basePath, contentsDirName)
	fileRepository := libs_file.NewFileDiskRepository(path)

	trBucketRepository := transfer_content.NewRepository(fileRepository)
	return contents.NewRepository(trBucketRepository)
}

// NewContentService creates a new content service
func NewContentService(basePath string) contents.Service {
	path := filepath.Join(basePath, contentsDirName)
	fileService := libs_file.NewFileDiskService(path)

	trContentService := transfer_content.NewService(fileService)
	return contents.NewService(trContentService)
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
