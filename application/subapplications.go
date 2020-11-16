package application

import (
	"github.com/xmn-services/buckets-network/application/chains"
	"github.com/xmn-services/buckets-network/application/miners"
	application_peers "github.com/xmn-services/buckets-network/application/peers"
	"github.com/xmn-services/buckets-network/application/storages"
)

type subApplications struct {
	peerApp    application_peers.Application
	chainApp   chains.Application
	storageApp storages.Application
	minerApp   miners.Application
}

func createSubApplicationa(
	peerApp application_peers.Application,
	chainApp chains.Application,
	storageApp storages.Application,
	minerApp miners.Application,
) SubApplications {
	out := subApplications{
		peerApp:    peerApp,
		chainApp:   chainApp,
		storageApp: storageApp,
		minerApp:   minerApp,
	}

	return &out
}

// Peers returns the peers application
func (app *subApplications) Peers() application_peers.Application {
	return app.peerApp
}

// Chain returns the chain application
func (app *subApplications) Chain() chains.Application {
	return app.chainApp
}

// Storage returns the storage application
func (app *subApplications) Storage() storages.Application {
	return app.storageApp
}

// Miner returns the miner application
func (app *subApplications) Miner() miners.Application {
	return app.minerApp
}
