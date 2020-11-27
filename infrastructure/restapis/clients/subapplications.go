package clients

import (
	"github.com/xmn-services/buckets-network/application/commands"
	"github.com/xmn-services/buckets-network/application/commands/chains"
	"github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/application/commands/storages"
)

type subApplications struct {
	peers   peers.Application
	chain   chains.Application
	storage storages.Application
}

func createSubApplications(
	peers peers.Application,
	chain chains.Application,
	storage storages.Application,
) commands.SubApplications {
	out := subApplications{
		peers:   peers,
		chain:   chain,
		storage: storage,
	}

	return &out
}

// Peers returns the peers application
func (obj *subApplications) Peers() peers.Application {
	return obj.peers
}

// Chain returns the chain application
func (obj *subApplications) Chain() chains.Application {
	return obj.chain
}

// Storage returns the storage application
func (obj *subApplications) Storage() storages.Application {
	return obj.storage
}
