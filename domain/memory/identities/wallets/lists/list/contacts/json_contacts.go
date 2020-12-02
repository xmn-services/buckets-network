package contacts

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"

// JSONContacts represents a JSON contact
type JSONContacts struct {
	List []*contact.JSONContact `json:"contacts"`
}

func createJSONContactsFromContacts(ins Contacts) *JSONContacts {
	adapter := contact.NewAdapter()
	lst := []*contact.JSONContact{}
	contactes := ins.All()
	for _, oneContact := range contactes {
		lst = append(lst, adapter.ToJSON(oneContact))
	}

	return createJSONContacts(
		lst,
	)
}

func createJSONContacts(
	lst []*contact.JSONContact,
) *JSONContacts {
	out := JSONContacts{
		List: lst,
	}

	return &out
}
