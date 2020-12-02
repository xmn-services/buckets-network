package accesses

import "github.com/xmn-services/buckets-network/libs/hash"

type builder struct {
	lst []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(lst []hash.Hash) Builder {
	app.lst = lst
	return app
}

// Now builds a new Accesses instance
func (app *builder) Now() (Accesses, error) {
	if app.lst == nil {
		app.lst = []hash.Hash{}
	}

	mp := map[string]hash.Hash{}
	for _, oneHash := range app.lst {
		mp[oneHash.String()] = oneHash
	}

	return createAccesses(app.lst, mp), nil
}
