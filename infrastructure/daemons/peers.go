package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application/identities/daemons"
	peers_app "github.com/xmn-services/buckets-network/application/peers"
	domain_peers "github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	client_peer "github.com/xmn-services/buckets-network/infrastructure/clients/peers"
)

type peers struct {
	remoteApplicationBuilder client_peer.Builder
	localApplication         peers_app.Application
	peersBuilder             domain_peers.Builder
	peersService             domain_peers.Service
	waitPeriod               time.Duration
	isStarted                bool
}

func createPeers(
	remoteApplicationBuilder client_peer.Builder,
	localApplication peers_app.Application,
	peersBuilder domain_peers.Builder,
	peersService domain_peers.Service,
	waitPeriod time.Duration,
) daemons.Application {
	out := peers{
		remoteApplicationBuilder: remoteApplicationBuilder,
		localApplication:         localApplication,
		peersBuilder:             peersBuilder,
		peersService:             peersService,
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

		peers, err := app.localApplication.Retrieve()
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

			remotePeers, err := remoteApplication.Retrieve()
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

		err = app.peersService.Save(updatedPeers)
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
