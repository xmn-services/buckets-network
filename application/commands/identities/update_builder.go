package identities

import "errors"

type updateBuilder struct {
	seed     string
	name     string
	password string
	root     string
}

func createUpdateBuilder() UpdateBuilder {
	out := updateBuilder{
		seed:     "",
		name:     "",
		password: "",
		root:     "",
	}

	return &out
}

// Create initializes the builder
func (app *updateBuilder) Create() UpdateBuilder {
	return createUpdateBuilder()
}

// WithSeed adds a seed to the builder
func (app *updateBuilder) WithSeed(seed string) UpdateBuilder {
	app.seed = seed
	return app
}

// WithName adds a name to the builder
func (app *updateBuilder) WithName(name string) UpdateBuilder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *updateBuilder) WithPassword(password string) UpdateBuilder {
	app.password = password
	return app
}

// WithRoot adds a root to the builder
func (app *updateBuilder) WithRoot(root string) UpdateBuilder {
	app.root = root
	return app
}

// Now builds a new Update instance
func (app *updateBuilder) Now() (Update, error) {
	if app.seed == "" && app.name == "" && app.password == "" && app.root == "" {
		return nil, errors.New("the Update instance cannot be empty")
	}

	return createUpdate(app.seed, app.name, app.password, app.root), nil
}
