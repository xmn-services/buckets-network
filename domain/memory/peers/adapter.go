package peers

import "encoding/json"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// JSONToPeers converts JSON to peers
func (app *adapter) JSONToPeers(js []byte) (Peers, error) {
	ins := new(peers)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
