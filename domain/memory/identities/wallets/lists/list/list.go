package list

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type list struct {
	hash        hash.Hash
	name        string
	description string
	contacts    contacts.Contacts
}

func createListFromJSON(ins *JSONList) (List, error) {
	contactsAdapter := contacts.NewAdapter()
	contacts, err := contactsAdapter.ToContacts(ins.Contacts)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithName(ins.Name).
		WithDescription(ins.Description).
		WithContacts(contacts).
		Now()
}

func createList(
	hash hash.Hash,
	name string,
	description string,
	contacts contacts.Contacts,
) List {
	out := list{
		hash:        hash,
		name:        name,
		description: description,
		contacts:    contacts,
	}

	return &out
}

// Hash returns the hash
func (obj *list) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *list) Name() string {
	return obj.name
}

// SetName sets the name
func (obj *list) SetName(name string) {
	obj.name = name
}

// Description returns the description
func (obj *list) Description() string {
	return obj.description
}

// SetDescription sets the description
func (obj *list) SetDescription(description string) {
	obj.description = description
}

// Contacts returns the contacts
func (obj *list) Contacts() contacts.Contacts {
	return obj.contacts
}

// MarshalJSON converts the instance to JSON
func (obj *list) MarshalJSON() ([]byte, error) {
	ins := createJSONListFromList(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *list) UnmarshalJSON(data []byte) error {
	ins := new(JSONList)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createListFromJSON(ins)
	if err != nil {
		return err
	}

	insList := pr.(*list)
	obj.hash = insList.hash
	obj.name = insList.name
	obj.description = insList.description
	obj.contacts = insList.contacts
	return nil
}
