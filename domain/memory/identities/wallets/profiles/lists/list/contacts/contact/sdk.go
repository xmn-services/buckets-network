package contact

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
