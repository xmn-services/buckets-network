package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	domain_peers "github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/clients"
)

type peers struct {
	remoteApplicationBuilder clients.Builder
	localApplication         application.Application
	peersBuilder             domain_peers.Builder
	waitPeriod               time.Duration
	isStarted                bool
}

func createPeers(
	remoteApplicationBuilder clients.Builder,
	localApplication application.Application,
	peersBuilder domain_peers.Builder,
	waitPeriod time.Duration,
) daemons.Application {
	out := peers{
		remoteApplicationBuilder: remoteApplicationBuilder,
		localApplication:         localApplication,
		peersBuilder:             peersBuilder,
		waitPeriod:               waitPeriod,
		isStarted:                false,
	}

	return &out
}

// Start starts the application
func (app *peers) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		peers, err := app.localApplication.Sub().Peers().Retrieve()
		if err != nil {
			return err
		}

		allPeers := []peer.Peer{}
		localPeers := peers.All()
		for _, oneLocalPeer := range localPeers {
			remoteApplication, err := app.remoteApplicationBuilder.Create().WithPeer(oneLocalPeer).Now()
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

		err = app.localApplication.Sub().Peers().Save(updatedPeers)
		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *peers) Stop() error {
	app.isStarted = true
	return nil
}
