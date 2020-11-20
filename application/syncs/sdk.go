package syncs

import (
	app "github.com/xmn-services/buckets-network/application/commands"
	domain_peer "github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// Builder represents a sync application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// ClientBuilder represents a client builder
type ClientBuilder interface {
	Create() ClientBuilder
	WithPeer(peer domain_peer.Peer) ClientBuilder
	Now() (app.Application, error)
}

// Application represents the sync application
type Application interface {
	Chain() error
	Peers() error
	Storage() error
	Miners() error
	MinerBlock() error
	MinerLink() error
}
