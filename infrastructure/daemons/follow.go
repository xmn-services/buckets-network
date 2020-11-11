package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application"
	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/infrastructure/clients"
)

type follow struct {
	application      app.Application
	remoteAppBuilder clients.Builder
	waitPeriod       time.Duration
	name             string
	password         string
	seed             string
	isStarted        bool
}

func createFollow(
	application app.Application,
	remoteAppBuilder clients.Builder,
	waitPeriod time.Duration,
	name string,
	password string,
	seed string,
	isStarted bool,
) daemons.Application {
	out := follow{
		application:      application,
		remoteAppBuilder: remoteAppBuilder,
		waitPeriod:       waitPeriod,
		name:             name,
		password:         password,
		seed:             seed,
		isStarted:        isStarted,
	}

	return &out
}

// Start starts the application
func (app *follow) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the identity:
		identityApp, err := app.application.Current().Authenticate(app.name, app.password, app.seed)
		if err != nil {
			return err
		}

		identity, err := identityApp.Current().Retrieve()
		if err != nil {
			return err
		}

		// retrieve the peers:
		peers, err := app.application.Sub().Peers().Retrieve()
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

		followBuckets := identity.Wallet().Follows().Requests()
		for _, oneBucketHash := range followBuckets {
			for _, oneRemoteApp := range remoteApps {
				bucket, err := oneRemoteApp.Sub().Follows().Retrieve(oneBucketHash)
				if err != nil {
					continue
				}

				err = identity.Wallet().Follows().Add(bucket)
				if err != nil {
					return err
				}

				break
			}
		}

		// save the identity:
		err = app.application.Current().UpdateIdentity(
			identity,
			app.password,
			app.password,
		)

		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *follow) Stop() error {
	app.isStarted = true
	return nil
}
