package contact

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact/accesses"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
)

// JSONContact represents a JSON access
type JSONContact struct {
	Hash        string                 `json:"hash"`
	Key         string                 `json:"pubkey"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Access      *accesses.JSONAccesses `json:"access"`
}

func createJSONContactFromContact(ins Contact) *JSONContact {
	accessAdapter := accesses.NewAdapter()
	access := accessAdapter.ToJSON(ins.Access())

	pubKeyAdapter := public.NewAdapter()
	pubKey := pubKeyAdapter.ToEncoded(ins.Key())
	hsh := ins.Hash().String()
	name := ins.Name()
	description := ins.Description()
	return createJSONContact(
		hsh,
		pubKey,
		name,
		description,
		access,
	)
}

func createJSONContact(
	hsh string,
	key string,
	name string,
	description string,
	access *accesses.JSONAccesses,
) *JSONContact {
	out := JSONContact{
		Hash:        hsh,
		Key:         key,
		Name:        name,
		Description: description,
		Access:      access,
	}

	return &out
}
