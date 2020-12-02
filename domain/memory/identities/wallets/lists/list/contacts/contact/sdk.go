package contact

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	pubKeyAdapter := public.NewAdapter()
	accessFactory := accesses.NewFactory()
	return createBuilder(hashAdapter, pubKeyAdapter, accessFactory)
}

// Adapter returns the contact adapter
type Adapter interface {
	ToJSON(contact Contact) *JSONContact
	ToContact(ins *JSONContact) (Contact, error)
}

// Builder represents a contact builder
type Builder interface {
	Create() Builder
	WithKey(key public.Key) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithAccess(access accesses.Accesses) Builder
	Now() (Contact, error)
}

// Contact represents a contact
type Contact interface {
	Hash() hash.Hash
	Key() public.Key
	Name() string
	Description() string
	Access() accesses.Accesses
}
