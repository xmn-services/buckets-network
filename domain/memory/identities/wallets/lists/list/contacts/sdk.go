package contacts

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter returns the contacts adapter
type Adapter interface {
	ToJSON(contacts Contacts) *JSONContacts
	ToContacts(ins *JSONContacts) (Contacts, error)
}

// Factory represents a contacts factory
type Factory interface {
	Create() (Contacts, error)
}

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithList(lst []contact.Contact) Builder
	Now() (Contacts, error)
}

// Contacts represents contacts
type Contacts interface {
	All() []contact.Contact
	Fetch(contactHash hash.Hash) (contact.Contact, error)
	Add(contact contact.Contact) error
	Delete(contact hash.Hash) error
	Update(contact contact.Contact) error
}
