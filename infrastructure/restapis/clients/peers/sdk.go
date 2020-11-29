package peers

import (
	"github.com/go-resty/resty/v2"
	commands_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application instance
func NewApplication(peerIns peer.Peer) commands_peers.Application {
	peersAdapter := peers.NewAdapter()
	peerAdapter := peer.NewAdapter()
	peerBuilder := peer.NewBuilder()
	client := resty.New()
	return createApplication(peersAdapter, peerAdapter, peerBuilder, client, peerIns)
}
