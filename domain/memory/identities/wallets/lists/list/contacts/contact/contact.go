package contact

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact/accesses"
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

func createContactFromJSON(ins *JSONContact) (Contact, error) {
	accessesAdapter := accesses.NewAdapter()
	accesses, err := accessesAdapter.ToAccesses(ins.Access)
	if err != nil {
		return nil, err
	}

	pubKeyAdapter := public.NewAdapter()
	pubKey, err := pubKeyAdapter.FromEncoded(ins.Key)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithKey(pubKey).
		WithName(ins.Name).
		WithDescription(ins.Description).
		WithAccess(accesses).
		Now()
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

// MarshalJSON converts the instance to JSON
func (obj *contact) MarshalJSON() ([]byte, error) {
	ins := createJSONContactFromContact(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *contact) UnmarshalJSON(data []byte) error {
	ins := new(JSONContact)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createContactFromJSON(ins)
	if err != nil {
		return err
	}

	insContact := pr.(*contact)
	obj.hash = insContact.hash
	obj.key = insContact.key
	obj.name = insContact.name
	obj.description = insContact.description
	obj.access = insContact.access
	return nil
}
