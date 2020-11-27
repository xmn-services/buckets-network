package buckets

import (
	"errors"

	"github.com/xmn-services/buckets-network/application/commands/identities/buckets"
)

type builder struct {
	name     string
	password string
	seed     string
}

func createBuilder() buckets.Builder {
	out := builder{
		name:     "",
		password: "",
		seed:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() buckets.Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) buckets.Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) buckets.Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) buckets.Builder {
	app.seed = seed
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (buckets.Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Application instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Application instance")
	}

	return nil, nil
}
