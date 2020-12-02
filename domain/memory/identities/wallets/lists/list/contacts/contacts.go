package contacts

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type contacts struct {
	lst []contact.Contact
	mp  map[string]contact.Contact
}

func createContactsFromJSON(ins *JSONContacts) (Contacts, error) {
	lst := []contact.Contact{}
	adapter := contact.NewAdapter()
	for _, oneContact := range ins.List {
		contact, err := adapter.ToContact(oneContact)
		if err != nil {
			return nil, err
		}

		lst = append(lst, contact)
	}

	return NewBuilder().
		Create().
		WithList(lst).
		Now()
}

func createContacts(
	lst []contact.Contact,
	mp map[string]contact.Contact,
) Contacts {
	out := contacts{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All returns the contacts
func (obj *contacts) All() []contact.Contact {
	return obj.lst
}

// Add adds a conatct
func (obj *contacts) Add(contact contact.Contact) error {
	keyname := contact.Hash().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the contact (hash: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, contact)
	obj.mp[keyname] = contact
	return nil
}

// Delete deletes a contact
func (obj *contacts) Delete(contact hash.Hash) error {
	keyname := contact.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the contact (hash: %s) does not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneContact := range obj.lst {
		if oneContact.Hash().Compare(contact) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *contacts) MarshalJSON() ([]byte, error) {
	ins := createJSONContactsFromContacts(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *contacts) UnmarshalJSON(data []byte) error {
	ins := new(JSONContacts)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createContactsFromJSON(ins)
	if err != nil {
		return err
	}

	insContacts := pr.(*contacts)
	obj.lst = insContacts.lst
	obj.mp = insContacts.mp
	return nil
}
