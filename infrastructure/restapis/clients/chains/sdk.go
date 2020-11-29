package chains

import (
	"github.com/go-resty/resty/v2"
	command_chains "github.com/xmn-services/buckets-network/application/commands/chains"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application instance
func NewApplication(peer peer.Peer) command_chains.Application {
	chainAdapter := chains.NewAdapter()
	client := resty.New()
	return createApplication(chainAdapter, client, peer)
}
