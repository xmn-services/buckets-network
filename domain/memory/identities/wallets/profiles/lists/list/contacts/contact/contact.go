package contact

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type contact struct {
	hash        hash.Hash
	key         public.Key
	name        string
	description string
	access      accesses.Accesses
}

func createContact(
	hash hash.Hash,
	key public.Key,
	name string,
	description string,
	access accesses.Accesses,
) Contact {
	out := contact{
		hash:        hash,
		key:         key,
		name:        name,
		description: description,
		access:      access,
	}

	return &out
}

// Hash returns the hash
func (obj *contact) Hash() hash.Hash {
	return obj.hash
}

// Key returns the key
func (obj *contact) Key() public.Key {
	return obj.key
}

// Name returns the name
func (obj *contact) Name() string {
	return obj.name
}

// Description returns the description
func (obj *contact) Description() string {
	return obj.description
}

// Access returns the access
func (obj *contact) Access() accesses.Accesses {
	return obj.access
}
