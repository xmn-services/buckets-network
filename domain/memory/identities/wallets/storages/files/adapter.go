package files

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts JSON to Files instance
func (app *adapter) ToJSON(ins Files) *JSONFiles {
	return createJSONFilesFromFiles(ins)
}

// ToFiles converts Files instance to JSON
func (app *adapter) ToFiles(js *JSONFiles) (Files, error) {
	return createFilesFromJSON(js)
}
