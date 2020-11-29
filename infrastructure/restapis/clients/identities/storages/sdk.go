package storages

import (
	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewBuilder creates a new storage builder
func NewBuilder(peer peer.Peer) storages.Builder {
	client := resty.New()
	return createBuilder(client, peer)
}
