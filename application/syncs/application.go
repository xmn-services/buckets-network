package syncs

import (
	"github.com/xmn-services/buckets-network/application/commands"
	application_chain "github.com/xmn-services/buckets-network/application/commands/chains"
	application_identity "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files/contents"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type application struct {
	appli                     commands.Application
	clientBuilder             commands.ClientBuilder
	chainBuilder              chains.Builder
	chainService              chains.Service
	peersBuilder              peers.Builder
	name                      string
	password                  string
	seed                      string
	additionalBucketsPerBlock uint
}

func createApplication(
	appli commands.Application,
	clientBuilder commands.ClientBuilder,
	chainBuilder chains.Builder,
	chainService chains.Service,
	peersBuilder peers.Builder,
	name string,
	password string,
	seed string,
	additionalBucketsPerBlock uint,
) Application {
	out := application{
		appli:                     appli,
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
	// authenticate:
	identityApp, err := app.appli.Current().Authenticate(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// mine the next block:
	err = identityApp.Sub().Chain().Block(app.additionalBucketsPerBlock)
	if err != nil {
		return err
	}

	// mine the next link:
	err = identityApp.Sub().Chain().Link(app.additionalBucketsPerBlock)
	if err != nil {
		return err
	}

	// retrieve the local chain:
	localChain, err := app.appli.Sub().Chain().Retrieve()
	if err != nil {
		return err
	}

	// retrieve the peers:
	localPeers, err := app.appli.Sub().Peers().Retrieve()
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
		gen := localChain.Genesis()
		root := localChain.Root()
		total := localChain.Total() + 1
		updatedChain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(remoteHead).WithTotal(total).Now()
		if err != nil {
			return err
		}

		err = app.chainService.Update(localChain, updatedChain)
		if err != nil {
			return err
		}
	}

	return nil
}

// Peers syncs the peers
func (app *application) Peers() error {
	peers, err := app.appli.Sub().Peers().Retrieve()
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

	updatedPeersList := updatedPeers.All()
	for _, onePeer := range updatedPeersList {
		err = app.appli.Sub().Peers().Save(onePeer)
		if err != nil {
			return err
		}
	}

	return nil
}

// Storage syncs the storage
func (app *application) Storage() error {
	identityApp, err := app.appli.Current().Authenticate(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// retrieve the identity:
	identity, err := identityApp.Current().Retrieve()
	if err != nil {
		return err
	}

	// retrieve the peers:
	peers, err := app.appli.Sub().Peers().Retrieve()
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
	toDownloadContents := identity.Wallet().Storage().ToDownload().All()
	err = app.download(identityApp, toDownloadContents, remoteApps)
	if err != nil {
		return err
	}

	// save the identity:
	return app.appli.Current().UpdateIdentity(identity, app.password, app.password)
}

func (app *application) download(
	identityApp application_identity.Application,
	toDownloadContents []contents.Content,
	clientApplication []commands.Application,
) error {
	for _, oneContent := range toDownloadContents {
		bucketHashStr := oneContent.Bucket().String()
		chunkHashStr := oneContent.Chunk().String()
		for _, oneClient := range clientApplication {
			clientStorageApp := oneClient.Sub().Storage()
			if !clientStorageApp.Exists(bucketHashStr, chunkHashStr) {
				continue
			}

			storedChunk, err := clientStorageApp.Retrieve(bucketHashStr, chunkHashStr)
			if err != nil {
				return err
			}

			// save the file:
			err = identityApp.Sub().Storage().Save(bucketHashStr, storedChunk)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
