package peers

import (
	application_peer "github.com/xmn-services/buckets-network/application/peers"
	domain_peer "github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// Builder represents a peer client builder
type Builder interface {
	Create() Builder
	WithPeer(peer domain_peer.Peer) Builder
	Now() (application_peer.Application, error)
}
