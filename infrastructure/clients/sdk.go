package clients

import (
	"github.com/xmn-services/buckets-network/application"
	domain_peer "github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// Builder represents a client builder
type Builder interface {
	Create() Builder
	WithPeer(peer domain_peer.Peer) Builder
	Now() (application.Application, error)
}
