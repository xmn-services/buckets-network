package chains

import (
	"encoding/json"

	transfer_chains "github.com/xmn-services/buckets-network/domain/transfers/chains"
)

type adapter struct {
	trBuilder transfer_chains.Builder
}

func createAdapter(
	trBuilder transfer_chains.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// JSONToChain converts a json to Chain instance
func (app *adapter) JSONToChain(js []byte) (Chain, error) {
	ins := new(chain)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// ToTransfer converts a chain to a transfer chain instance
func (app *adapter) ToTransfer(chain Chain) (transfer_chains.Chain, error) {
	hash := chain.Hash()
	gen := chain.Genesis().Hash()
	root := chain.Root().Hash()
	head := chain.Head().Hash()
	total := chain.Total()
	createdOn := chain.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hash).
		WithGenesis(gen).
		WithRoot(root).
		WithHead(head).
		WithTotal(total).
		CreatedOn(createdOn).
		Now()
}
