package lists

import "errors"

type updateBuilder struct {
	name        string
	description string
}

func createUpdateBuilder() UpdateBuilder {
	out := updateBuilder{
		name:        "",
		description: "",
	}

	return &out
}

// Create initializes the builder
func (app *updateBuilder) Create() UpdateBuilder {
	return createUpdateBuilder()
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

// Now updates a new Update instance
func (app *updateBuilder) Now() (Update, error) {
	if app.name != "" && app.description != "" {
		return createUpdateWithNameAndDescription(app.name, app.description), nil
	}

	if app.name != "" {
		return createUpdateWithName(app.name), nil
	}

	if app.description != "" {
		return createUpdateWithDescription(app.description), nil
	}

	return nil, errors.New("the Update is invalid")
}
