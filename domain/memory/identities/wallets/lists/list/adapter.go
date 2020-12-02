package list

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an list to JSON
func (app *adapter) ToJSON(list List) *JSONList {
	return createJSONListFromList(list)
}

// ToList converts JSON to list
func (app *adapter) ToList(ins *JSONList) (List, error) {
	return createListFromJSON(ins)
}
