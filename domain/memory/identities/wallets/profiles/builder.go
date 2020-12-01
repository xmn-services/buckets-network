package profiles

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists"
)

type builder struct {
	listFactory   lists.Factory
	accessFactory accesses.Factory
	name          string
	description   string
	list          lists.Lists
	access        accesses.Accesses
}

func createBuilder(
	listFactory lists.Factory,
	accessFactory accesses.Factory,
) Builder {
	out := builder{
		listFactory:   listFactory,
		accessFactory: accessFactory,
		name:          "",
		description:   "",
		list:          nil,
		access:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.listFactory, app.accessFactory)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithLists adds a lists to the builder
func (app *builder) WithLists(list lists.Lists) Builder {
	app.list = list
	return app
}

// WitWithAccesshLists adds an access to the builder
func (app *builder) WithAccess(access accesses.Accesses) Builder {
	app.access = access
	return app
}

// Now builds a new Profile instance
func (app *builder) Now() (Profile, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Profile instance")
	}

	if app.list == nil {
		app.list = app.listFactory.Create()
	}

	if app.access == nil {
		app.access = app.accessFactory.Create()
	}

	return createProfile(app.name, app.description, app.list, app.access), nil
}
