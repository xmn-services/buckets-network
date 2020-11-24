package peer

import (
	"net/url"

	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents a peer adapter
type Adapter interface {
	URLValuesToPeer(values url.Values) (Peer, error)
	PeerToURLValues(peer Peer) (url.Values, error)
	StringToPeer(str string) (Peer, error)
}

// Builder represents a peer builder
type Builder interface {
	Create() Builder
	WithHost(host string) Builder
	WithPort(port uint) Builder
	IsClear() Builder
	IsOnion() Builder
	Now() (Peer, error)
}

// Peer represents a peer
type Peer interface {
	Host() string
	Port() uint
	IsClear() bool
	IsOnion() bool
	String() string
}
