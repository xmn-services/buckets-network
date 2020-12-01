package list

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type list struct {
	hash        hash.Hash
	name        string
	description string
	contacts    contacts.Contacts
}

func createList(
	hash hash.Hash,
	name string,
	description string,
	contacts contacts.Contacts,
) List {
	out := list{
		hash:        hash,
		name:        name,
		description: description,
		contacts:    contacts,
	}

	return &out
}

// Hash returns the hash
func (obj *list) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *list) Name() string {
	return obj.name
}

// Description returns the description
func (obj *list) Description() string {
	return obj.description
}

// Contacts returns the contacts
func (obj *list) Contacts() contacts.Contacts {
	return obj.contacts
}
