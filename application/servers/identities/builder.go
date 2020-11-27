package identities

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/servers/authenticates"
)

type builder struct {
	auth authenticates.Authenticate
	root string
}

func createBuilder() Builder {
	out := builder{
		auth: nil,
		root: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAuthenticate adds an authenticate instance to the builder
func (app *builder) WithAuthenticate(auth authenticates.Authenticate) Builder {
	app.auth = auth
	return app
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.auth == nil {
		return nil, errors.New("the Authenticate instance is mandatory in order to build an Identity instance")
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build an Identity instance")
	}

	return createIdentity(app.auth, app.root), nil
}
