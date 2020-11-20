package peer

import "github.com/xmn-services/buckets-network/libs/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
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
