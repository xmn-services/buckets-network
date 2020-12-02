package contact

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an contact to JSON
func (app *adapter) ToJSON(contact Contact) *JSONContact {
	return createJSONContactFromContact(contact)
}

// ToContact converts JSON to contact
func (app *adapter) ToContact(ins *JSONContact) (Contact, error) {
	return createContactFromJSON(ins)
}
