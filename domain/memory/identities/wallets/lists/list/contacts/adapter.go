package contacts

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an accesses to JSON
func (app *adapter) ToJSON(accesses Contacts) *JSONContacts {
	return createJSONContactsFromContacts(accesses)
}

// ToContacts converts JSON to accesses
func (app *adapter) ToContacts(ins *JSONContacts) (Contacts, error) {
	return createContactsFromJSON(ins)
}
