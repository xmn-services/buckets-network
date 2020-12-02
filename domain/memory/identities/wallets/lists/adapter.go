package lists

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an lists to JSON
func (app *adapter) ToJSON(lists Lists) *JSONLists {
	return createJSONListsFromLists(lists)
}

// ToLists converts JSON to lists
func (app *adapter) ToLists(ins *JSONLists) (Lists, error) {
	return createListsFromJSON(ins)
}
