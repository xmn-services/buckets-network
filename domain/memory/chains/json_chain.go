package chains

import (
	"time"

	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

// JSONChain represents a json chain
type JSONChain struct {
	Genesis   *genesis.JSONGenesis   `json:"genesis"`
	Root      *mined_block.JSONBlock `json:"root"`
	Head      *mined_link.JSONLink   `json:"head"`
	Total     uint                   `json:"total"`
	CreatedOn time.Time              `json:"created_on"`
}

func createJSONChainFromChain(ins Chain) *JSONChain {
	genesisAdapter := genesis.NewAdapter()
	genesis := genesisAdapter.ToJSON(ins.Genesis())

	blockAdapter := mined_block.NewAdapter()
	root := blockAdapter.ToJSON(ins.Root())

	linkAdapter := mined_link.NewAdapter()
	head := linkAdapter.ToJSON(ins.Head())

	total := ins.Total()
	createdOn := ins.CreatedOn()
	return createJSONChain(genesis, root, head, total, createdOn)
}

func createJSONChain(
	genesis *genesis.JSONGenesis,
	root *mined_block.JSONBlock,
	head *mined_link.JSONLink,
	total uint,
	createdOn time.Time,
) *JSONChain {
	out := JSONChain{
		Genesis:   genesis,
		Root:      root,
		Head:      head,
		Total:     total,
		CreatedOn: createdOn,
	}

	return &out
}
