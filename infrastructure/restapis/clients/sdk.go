package clients

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands"
	"github.com/xmn-services/buckets-network/application/commands/chains"
	command_identities "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/application/commands/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application instance
func NewApplication(
	commandIdentityBuilder command_identities.Builder,
	peers peers.Application,
	chain chains.Application,
	storage storages.Application,
	peer peer.Peer,
) commands.Application {
	client := resty.New()
	subApplications := createSubApplications(peers, chain, storage)
	current := createCurrent(commandIdentityBuilder, client, peer)
	return createApplication(current, subApplications)
}
