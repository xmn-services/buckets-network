package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application"
	identities_app "github.com/xmn-services/buckets-network/application/identities"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/infrastructure/clients"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type storage struct {
	localApplication application.Application
	remoteAppBuilder clients.Builder
	name             string
	seed             string
	password         string
	waitPeriod       time.Duration
	isStarted        bool
}

func createStorage(
	localApplication application.Application,
	remoteAppBuilder clients.Builder,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
) daemons.Application {
	out := storage{
		localApplication: localApplication,
		remoteAppBuilder: remoteAppBuilder,
		waitPeriod:       waitPeriod,
		name:             name,
		seed:             seed,
		password:         password,
		isStarted:        false,
	}

	return &out
}

// Start starts the application
func (app *storage) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the identity:
		identityApp, err := app.localApplication.Current().Authenticate(app.name, app.password, app.seed)
		if err != nil {
			return err
		}

		identity, err := identityApp.Current().Retrieve()
		if err != nil {
			return err
		}

		// retrieve the peers:
		peers, err := app.localApplication.Sub().Peers().Retrieve()
		if err != nil {
			return err
		}

		// build the remote applications:
		allPeers := peers.All()
		remoteApps := []application.Application{}
		for _, onePeer := range allPeers {
			remoteApp, err := app.remoteAppBuilder.Create().WithPeer(onePeer).Now()
			if err != nil {
				return err
			}

			remoteApps = append(remoteApps, remoteApp)
		}

		// download the files:
		toDownloadFiles := identity.Wallet().Storage().ToDownload().All()
		err = app.download(toDownloadFiles, identityApp, remoteApps)
		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *storage) Stop() error {
	app.isStarted = true
	return nil
}

func (app *storage) download(toDownloadFiles []hash.Hash, identityApp identities_app.Application, clientApplication []application.Application) error {
	for _, oneFileHash := range toDownloadFiles {
		for _, oneClient := range clientApplication {
			clientStorageApp := oneClient.Sub().Storage()
			if !clientStorageApp.IsStored(oneFileHash) {
				continue
			}

			storedFile, err := clientStorageApp.Retrieve(oneFileHash)
			if err != nil {
				return err
			}

			// save the file:
			err = identityApp.Sub().Storage().Save(storedFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
