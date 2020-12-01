package list

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts"
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
