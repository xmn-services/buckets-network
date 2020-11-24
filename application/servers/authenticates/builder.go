package authenticates

import "errors"

type builder struct {
	name     string
	password string
	seed     string
}

func createBuilder() Builder {
	out := builder{
		name:     "",
		password: "",
		seed:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// Now builds a new Authenticate instance
func (app *builder) Now() (Authenticate, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Authenticate instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build an Authenticate instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Authenticate instance")
	}

	return createAuthenticate(app.name, app.password, app.seed), nil
}
