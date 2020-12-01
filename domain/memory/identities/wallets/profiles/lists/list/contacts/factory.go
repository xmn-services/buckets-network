package contacts

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts/contact"

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new contacts instance
func (app *factory) Create() Contacts {
	lst := []contact.Contact{}
	mp := map[string]contact.Contact{}
	return createContacts(lst, mp)
}
