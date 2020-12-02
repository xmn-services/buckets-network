package list

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	contactsFactory := contacts.NewFactory()
	return createBuilder(hashAdapter, contactsFactory)
}

// Adapter returns the list adapter
type Adapter interface {
	ToJSON(list List) *JSONList
	ToList(ins *JSONList) (List, error)
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithContacts(contacts contacts.Contacts) Builder
	Now() (List, error)
}

// List represents a contact list
type List interface {
	Hash() hash.Hash
	Name() string
	Description() string
	Contacts() contacts.Contacts
}
