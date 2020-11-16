package miners

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts a miner to JSON
func (app *adapter) ToJSON(ins Miner) *JSONMiner {
	return createJSONMinerFromMiner(ins)
}

// ToMiner converts JSON to a Miner instance
func (app *adapter) ToMiner(js *JSONMiner) (Miner, error) {
	return createMinerFromJSON(js)
}
