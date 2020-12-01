package contacts

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a contacts factory
type Factory interface {
	Create() Contacts
}

// Contacts represents contacts
type Contacts interface {
	All() []contact.Contact
	Add(contact contact.Contact) error
	Delete(contact hash.Hash) error
}
