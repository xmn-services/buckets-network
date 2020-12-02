package list

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	contactsFactory contacts.Factory
	name            string
	description     string
	contacts        contacts.Contacts
}

func createBuilder(
	hashAdapter hash.Adapter,
	contactsFactory contacts.Factory,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		contactsFactory: contactsFactory,
		name:            "",
		description:     "",
		contacts:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.contactsFactory)
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

// WithContacts adds a contacts to the builder
func (app *builder) WithContacts(contacts contacts.Contacts) Builder {
	app.contacts = contacts
	return app
}

// Now builds a new List instance
func (app *builder) Now() (List, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a List instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte(app.description),
	})

	if err != nil {
		return nil, err
	}

	if app.contacts == nil {
		contacts, err := app.contactsFactory.Create()
		if err != nil {
			return nil, err
		}

		app.contacts = contacts
	}

	return createList(*hsh, app.name, app.description, app.contacts), nil
}
