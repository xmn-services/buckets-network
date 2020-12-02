package contacts

import (
	"errors"

	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
)

type updateBuilder struct {
	key         public.Key
	name        string
	description string
}

func createUpdateBuilder() UpdateBuilder {
	out := updateBuilder{
		key:         nil,
		name:        "",
		description: "",
	}

	return &out
}

// Create initializes the builder
func (app *updateBuilder) Create() UpdateBuilder {
	return createUpdateBuilder()
}

// WithKey adds a key to the builder
func (app *updateBuilder) WithKey(key public.Key) UpdateBuilder {
	app.key = key
	return app
}

// WithName adds a name to the builder
func (app *updateBuilder) WithName(name string) UpdateBuilder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *updateBuilder) WithDescription(description string) UpdateBuilder {
	app.description = description
	return app
}

// Now builds a new Update instance
func (app *updateBuilder) Now() (Update, error) {
	if app.key != nil && app.name != "" && app.description != "" {
		return createUpdateWithKeyAndNameAndDescription(app.key, app.name, app.description), nil
	}

	if app.key != nil && app.name != "" {
		return createUpdateWithKeyAndName(app.key, app.name), nil
	}

	if app.key != nil && app.description != "" {
		return createUpdateWithKeyAndDescription(app.key, app.description), nil
	}

	if app.name != "" && app.description != "" {
		return createUpdateWithNameAndDescription(app.name, app.description), nil
	}

	if app.key != nil {
		return createUpdateWithKey(app.key), nil
	}

	if app.name != "" {
		return createUpdateWithName(app.name), nil
	}

	if app.description != "" {
		return createUpdateWithDescription(app.description), nil
	}

	return nil, errors.New("the Update is invalid")
}
