package lists

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"

type builder struct {
	lst []list.List
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
func (app *builder) WithList(lst []list.List) Builder {
	app.lst = lst
	return app
}

// Now builds a new Lists instance
func (app *builder) Now() (Lists, error) {
	if app.lst == nil {
		app.lst = []list.List{}
	}

	mp := map[string]list.List{}
	for _, oneList := range app.lst {
		mp[oneList.Hash().String()] = oneList
	}

	return createLists(app.lst, mp), nil
}
