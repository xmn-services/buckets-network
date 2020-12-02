package accesses

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"

type builder struct {
	lst []access.Access
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
func (app *builder) WithList(lst []access.Access) Builder {
	app.lst = lst
	return app
}

// Now builds a new Accesses instance
func (app *builder) Now() (Accesses, error) {
	if app.lst == nil {
		app.lst = []access.Access{}
	}

	mp := map[string]access.Access{}
	for _, oneAccess := range app.lst {
		mp[oneAccess.Bucket().String()] = oneAccess
	}

	return createAccesses(app.lst, mp), nil
}
