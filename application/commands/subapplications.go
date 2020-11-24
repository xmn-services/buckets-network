package commands

import (
	"github.com/xmn-services/buckets-network/application/commands/chains"
	application_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/application/commands/storages"
)

type subApplications struct {
	peerApp    application_peers.Application
	chainApp   chains.Application
	storageApp storages.Application
}

func createSubApplicationa(
	peerApp application_peers.Application,
	chainApp chains.Application,
	storageApp storages.Application,
) SubApplications {
	out := subApplications{
		peerApp:    peerApp,
		chainApp:   chainApp,
		storageApp: storageApp,
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
