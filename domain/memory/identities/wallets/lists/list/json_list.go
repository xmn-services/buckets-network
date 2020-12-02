package list

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"

// JSONList represents a JSON list
type JSONList struct {
	Hash        string                 `json:"hash"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Contacts    *contacts.JSONContacts `json:"contacts"`
}

func createJSONListFromList(ins List) *JSONList {
	contactAdapter := contacts.NewAdapter()
	contacts := contactAdapter.ToJSON(ins.Contacts())

	hsh := ins.Hash().String()
	name := ins.Name()
	description := ins.Description()
	return createJSONList(
		hsh,
		name,
		description,
		contacts,
	)
}

func createJSONList(
	hsh string,
	name string,
	description string,
	contacts *contacts.JSONContacts,
) *JSONList {
	out := JSONList{
		Hash:        hsh,
		Name:        name,
		Description: description,
		Contacts:    contacts,
	}

	return &out
}
