package contact

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	pubKeyAdapter public.Adapter
	accessFactory accesses.Factory
	key           public.Key
	name          string
	description   string
	access        accesses.Accesses
}

func createBuilder(
	hashAdapter hash.Adapter,
	pubKeyAdapter public.Adapter,
	accessFactory accesses.Factory,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		pubKeyAdapter: pubKeyAdapter,
		accessFactory: accessFactory,
		key:           nil,
		name:          "",
		description:   "",
		access:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.pubKeyAdapter,
		app.accessFactory,
	)
}

// WithKey addas a key to the builder
func (app *builder) WithKey(key public.Key) Builder {
	app.key = key
	return app
}

// WithName addas a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription addas a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithAccess adds an access to the builder
func (app *builder) WithAccess(access accesses.Accesses) Builder {
	app.access = access
	return app
}

// Now builds a new Contact instance
func (app *builder) Now() (Contact, error) {
	if app.key == nil {
		return nil, errors.New("the PublicKey is mandatory in order to build a Contact instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Contact instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.pubKeyAdapter.ToBytes(app.key),
		[]byte(app.name),
		[]byte(app.description),
	})

	if err != nil {
		return nil, err
	}

	if app.access == nil {
		app.access = app.accessFactory.Create()
	}

	return createContact(*hsh, app.key, app.name, app.description, app.access), nil
}
