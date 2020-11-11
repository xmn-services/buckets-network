package blocks

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts a blocks instance to JSON
func (app *adapter) ToJSON(ins Blocks) *JSONBlocks {
	return createJSONBlocksFromBlocks(ins)
}

// ToBlocks converts JSON to blocks instance
func (app *adapter) ToBlocks(js *JSONBlocks) (Blocks, error) {
	return createBlocksFromJSON(js)
}
