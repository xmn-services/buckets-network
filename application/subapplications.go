package application

import (
	"github.com/xmn-services/buckets-network/application/chains"
	"github.com/xmn-services/buckets-network/application/genesis"
	application_peers "github.com/xmn-services/buckets-network/application/peers"
)

type subApplications struct {
	peerApp    application_peers.Application
	genesisApp genesis.Application
	chainApp   chains.Application
}

func createSubApplicationa(
	peerApp application_peers.Application,
	genesisApp genesis.Application,
	chainApp chains.Application,
) SubApplications {
	out := subApplications{
		peerApp:    peerApp,
		genesisApp: genesisApp,
		chainApp:   chainApp,
	}

	return &out
}

// Peers returns the peers application
func (app *subApplications) Peers() application_peers.Application {
	return app.peerApp
}

// Genesis returns the genesis application
func (app *subApplications) Genesis() genesis.Application {
	return app.genesisApp
}

// Chain returns the chain application
func (app *subApplications) Chain() chains.Application {
	return app.chainApp
}
