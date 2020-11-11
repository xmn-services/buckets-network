package application

import (
	"github.com/xmn-services/buckets-network/application/chains"
	"github.com/xmn-services/buckets-network/application/follows"
	"github.com/xmn-services/buckets-network/application/genesis"
	application_identities "github.com/xmn-services/buckets-network/application/identities"
	application_peers "github.com/xmn-services/buckets-network/application/peers"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
)

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
	Genesis() genesis.Application
	Chain() chains.Application
	Follows() follows.Application
}
