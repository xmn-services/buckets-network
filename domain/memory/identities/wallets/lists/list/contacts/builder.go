package contacts

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"

type builder struct {
	lst []contact.Contact
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
func (app *builder) WithList(lst []contact.Contact) Builder {
	app.lst = lst
	return app
}

// Now builds a new Contacts instance
func (app *builder) Now() (Contacts, error) {
	if app.lst == nil {
		app.lst = []contact.Contact{}
	}

	mp := map[string]contact.Contact{}
	for _, oneAccess := range app.lst {
		mp[oneAccess.Hash().String()] = oneAccess
	}

	return createContacts(app.lst, mp), nil
}
