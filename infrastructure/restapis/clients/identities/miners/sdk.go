package miners

import (
	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewBuilder creates a new builder application
func NewBuilder(peer peer.Peer) miners.Builder {
	client := resty.New()
	return createBuilder(client, peer)
}
