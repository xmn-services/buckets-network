package storages

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts JSON to Storage instance
func (app *adapter) ToJSON(ins Storage) *JSONStorage {
	return createJSONStorageFromStorage(ins)
}

// ToStorage converts JSON to Storage instance
func (app *adapter) ToStorage(js *JSONStorage) (Storage, error) {
	return createStorageFromJSON(js)
}
