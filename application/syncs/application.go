package syncs

import (
	"github.com/xmn-services/buckets-network/application/commands"
	application_chain "github.com/xmn-services/buckets-network/application/commands/chains"
	application_identity_buckets "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	application_identity_storages "github.com/xmn-services/buckets-network/application/commands/identities/storages"
	application_storages "github.com/xmn-services/buckets-network/application/commands/identities/storages"
	application_miners "github.com/xmn-services/buckets-network/application/commands/miners"
	application_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	mined_links "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	chainApp                  application_chain.Application
	minerApp                  application_miners.Application
	peersApp                  application_peers.Application
	storageApp                application_storages.Application
	identityBucketApp         application_identity_buckets.Application
	identityStorageApp        application_identity_storages.Application
	identityRepository        identities.Repository
	identityService           identities.Service
	clientBuilder             ClientBuilder
	chainBuilder              chains.Builder
	chainService              chains.Service
	peersBuilder              peers.Builder
	name                      string
	password                  string
	seed                      string
	additionalBucketsPerBlock uint
}

func createApplication(
	chainApp application_chain.Application,
	minerApp application_miners.Application,
	peersApp application_peers.Application,
	storageApp application_storages.Application,
	identityBucketApp application_identity_buckets.Application,
	identityStorageApp application_identity_storages.Application,
	identityRepository identities.Repository,
	identityService identities.Service,
	clientBuilder ClientBuilder,
	chainBuilder chains.Builder,
	chainService chains.Service,
	peersBuilder peers.Builder,
	name string,
	password string,
	seed string,
	additionalBucketsPerBlock uint,
) Application {
	out := application{
		chainApp:                  chainApp,
		minerApp:                  minerApp,
		peersApp:                  peersApp,
		storageApp:                storageApp,
		identityBucketApp:         identityBucketApp,
		identityStorageApp:        identityStorageApp,
		identityRepository:        identityRepository,
		identityService:           identityService,
		clientBuilder:             clientBuilder,
		chainBuilder:              chainBuilder,
		chainService:              chainService,
		peersBuilder:              peersBuilder,
		name:                      name,
		password:                  password,
		seed:                      seed,
		additionalBucketsPerBlock: additionalBucketsPerBlock,
	}

	return &out
}

// Chain syncs the chain
func (app *application) Chain() error {
	// retrieve the local chain:
	localChain, err := app.chainApp.Retrieve()
	if err != nil {
		return err
	}

	// retrieve the peers:
	localPeers, err := app.peersApp.Retrieve()
	if err != nil {
		return err
	}

	biggestDiff := 0
	var biggestChain chains.Chain
	var biggestChainApp application_chain.Application
	allPeers := localPeers.All()
	for _, onePeer := range allPeers {
		remoteApp, err := app.clientBuilder.Create().WithPeer(onePeer).Now()
		if err != nil {
			return err
		}

		remoteChainApp := remoteApp.Sub().Chain()
		remoteChain, err := remoteChainApp.Retrieve()
		if err != nil {
			return err
		}

		diffTrx := int(remoteChain.Total() - localChain.Total())
		if biggestDiff < diffTrx {
			biggestDiff = diffTrx
			biggestChain = remoteChain
			biggestChainApp = remoteChainApp
		}
	}

	// if there is no chain in the network more advanced, return:
	if biggestChain == nil {
		return nil
	}

	// update the chain:
	localIndex := int(localChain.Height())
	diffHeight := int(biggestChain.Height()) - localIndex
	for i := 0; i < diffHeight; i++ {
		chainIndex := localIndex + i
		remoteChainAtIndex, err := biggestChainApp.RetrieveAtIndex(uint(chainIndex))
		if err != nil {
			return err
		}

		remoteHead := remoteChainAtIndex.Head()
		err = app.updateChain(localChain, remoteHead)
		if err != nil {
			return err
		}
	}

	return nil
}

// Peers syncs the peers
func (app *application) Peers() error {
	peers, err := app.peersApp.Retrieve()
	if err != nil {
		return err
	}

	allPeers := []peer.Peer{}
	localPeers := peers.All()
	for _, oneLocalPeer := range localPeers {
		remoteApplication, err := app.clientBuilder.Create().WithPeer(oneLocalPeer).Now()
		if err != nil {
			return err
		}

		remotePeers, err := remoteApplication.Sub().Peers().Retrieve()
		if err != nil {
			return err
		}

		allPeers = append(allPeers, oneLocalPeer)
		allPeers = append(allPeers, remotePeers.All()...)
	}

	updatedPeers, err := app.peersBuilder.Create().WithPeers(allPeers).Now()
	if err != nil {
		return err
	}

	return app.peersApp.Save(updatedPeers)
}

// Storage syncs the storage
func (app *application) Storage() error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve the peers:
	peers, err := app.peersApp.Retrieve()
	if err != nil {
		return err
	}

	// build the remote applications:
	allPeers := peers.All()
	remoteApps := []commands.Application{}
	for _, onePeer := range allPeers {
		remoteApp, err := app.clientBuilder.Create().WithPeer(onePeer).Now()
		if err != nil {
			return err
		}

		remoteApps = append(remoteApps, remoteApp)
	}

	// download the files:
	toDownloadFiles := identity.Wallet().Storage().ToDownload().All()
	err = app.download(toDownloadFiles, remoteApps)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Miners syncs the miners
func (app *application) Miners() error {
	err := app.MinerBlock()
	if err != nil {
		return err
	}

	return app.MinerLink()
}

// MinerBlock syncs the block miner
func (app *application) MinerBlock() error {
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve the queue buckets:
	queueBuckets := identity.Wallet().Miner().Queue().All()

	// make the list of block buckets:
	bucketHashes := []string{}
	for _, oneQueuedBucket := range queueBuckets {
		// add the bucket to the block list:
		bucketHashes = append(bucketHashes, oneQueuedBucket.Bucket().Hash().String())

		// remove the queued bucket from the queued identity:
		identity.Wallet().Miner().Broadcasted().Add(oneQueuedBucket)
	}

	// retrieve the chain:
	chain, err := app.chainApp.Retrieve()
	if err != nil {
		return err
	}

	blockDiff := chain.Genesis().Difficulty().Block()
	minedBlock, err := app.minerApp.Block(
		bucketHashes,
		blockDiff.Base(),
		blockDiff.IncreasePerBucket(),
		app.additionalBucketsPerBlock,
	)

	if err != nil {
		return err
	}

	// add the block to the wallet:
	err = identity.Wallet().Miner().ToLink().Add(minedBlock)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// MinerLink syncs the link miner
func (app *application) MinerLink() error {
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// retrieve the chain:
	chain, err := app.chainApp.Retrieve()
	if err != nil {
		return err
	}

	// fetch the mined blocks to link:
	blocks := identity.Wallet().Miner().ToLink().All()
	for _, oneBlock := range blocks {
		// mine the link:
		difficulty := chain.Genesis().Difficulty().Link()
		newMinedLink, err := app.minerApp.Link(
			chain.Head().Hash().String(),
			oneBlock.Hash().String(),
			difficulty,
		)

		if err != nil {
			return err
		}

		err = app.updateChain(chain, newMinedLink)
		if err != nil {
			return err
		}
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

func (app *application) updateChain(chain chains.Chain, newMinedLink mined_links.Link) error {
	gen := chain.Genesis()
	root := chain.Root()
	total := chain.Total() + 1
	updatedChain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(newMinedLink).WithTotal(total).Now()
	if err != nil {
		return err
	}

	return app.chainService.Update(chain, updatedChain)
}

func (app *application) download(toDownloadFiles []hash.Hash, clientApplication []commands.Application) error {
	for _, oneFileHash := range toDownloadFiles {
		for _, oneClient := range clientApplication {
			clientStorageApp := oneClient.Sub().Storage()
			if !clientStorageApp.IsStored(oneFileHash.String()) {
				continue
			}

			storedFile, err := clientStorageApp.Retrieve(oneFileHash.String())
			if err != nil {
				return err
			}

			// save the file:
			err = app.storageApp.Save(storedFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
