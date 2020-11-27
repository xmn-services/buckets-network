package identities

import (
	"github.com/xmn-services/buckets-network/application/servers/authenticates"
)

type identity struct {
	auth authenticates.Authenticate
	root string
}

func createIdentity(
	auth authenticates.Authenticate,
	root string,
) Identity {
	out := identity{
		auth: auth,
		root: root,
	}

	return &out
}

// Authenticate returns the authenticate instance
func (obj *identity) Authenticate() authenticates.Authenticate {
	return obj.auth
}

// Root returns the root
func (obj *identity) Root() string {
	return obj.root
}
