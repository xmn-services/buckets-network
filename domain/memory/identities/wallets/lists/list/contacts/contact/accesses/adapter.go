package accesses

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts an accesses to JSON
func (app *adapter) ToJSON(accesses Accesses) *JSONAccesses {
	return createJSONAccessesFromAccesses(accesses)
}

// ToAccesses converts JSON to accesses
func (app *adapter) ToAccesses(ins *JSONAccesses) (Accesses, error) {
	return createAccessesFromJSON(ins)
}
