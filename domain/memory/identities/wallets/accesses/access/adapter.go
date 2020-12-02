package access

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an access to JSON
func (app *adapter) ToJSON(access Access) *JSONAccess {
	return createJSONAccessFromAccess(access)
}

// ToAccess converts JSON to access
func (app *adapter) ToAccess(ins *JSONAccess) (Access, error) {
	return createAccessFromJSON(ins)
}
