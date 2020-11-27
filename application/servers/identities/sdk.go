package identities

import (
	"net/url"

	"github.com/xmn-services/buckets-network/application/servers/authenticates"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	authAdapter := authenticates.NewAdapter()
	builder := NewBuilder()
	return createAdapter(authAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the identity adapter
type Adapter interface {
	URLValuesToIdentity(urlValues url.Values) (Identity, error)
	IdentityToURLValues(identity Identity) url.Values
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithAuthenticate(auth authenticates.Authenticate) Builder
	WithRoot(root string) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Authenticate() authenticates.Authenticate
	Root() string
}
