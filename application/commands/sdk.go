package commands

import (
	"github.com/xmn-services/buckets-network/application/commands/chains"
	application_identities "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/application/commands/miners"
	application_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/application/commands/storages"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	domain_peer "github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

// NewApplication creates a new application instance
func NewApplication(
	peerApp application_peers.Application,
	chainApp chains.Application,
	storageApp storages.Application,
	minerApp miners.Application,
	identityAppBuilder application_identities.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
) Application {
	identityBuilder := identities.NewBuilder()
	current := createCurrent(identityAppBuilder, identityBuilder, identityRepository, identityService)
	subApps := createSubApplicationa(peerApp, chainApp, storageApp, minerApp)
	return createApplication(current, subApps)
}

// ClientBuilder represents a client builder
type ClientBuilder interface {
	Create() ClientBuilder
	WithPeer(peer domain_peer.Peer) ClientBuilder
	Now() (Application, error)
}

// Application represents the application
type Application interface {
	Current() Current
	Sub() SubApplications
}

// Current represents the current application
type Current interface {
	NewIdentity(name string, password string, seed string, root string) error
	Authenticate(name string, seed string, password string) (application_identities.Application, error)
	UpdateIdentity(identity identities.Identity, password string, newPassword string) error
}

// SubApplications represents the sub applications
type SubApplications interface {
	Peers() application_peers.Application
	Chain() chains.Application
	Storage() storages.Application
	Miner() miners.Application
}
