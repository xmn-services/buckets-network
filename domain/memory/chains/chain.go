package chains

import (
	"encoding/json"
	"time"

	mined_block "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type chain struct {
	immutable entities.Immutable
	genesis   genesis.Genesis
	root      mined_block.Block
	head      mined_link.Link
	total     uint
}

func createChainFromJSON(ins *JSONChain) (Chain, error) {
	genesisAdapter := genesis.NewAdapter()
	genesis, err := genesisAdapter.ToGenesis(ins.Genesis)
	if err != nil {
		return nil, err
	}

	blockAdapter := mined_block.NewAdapter()
	root, err := blockAdapter.ToBlock(ins.Root)
	if err != nil {
		return nil, err
	}

	linkAdapter := mined_link.NewAdapter()
	head, err := linkAdapter.ToLink(ins.Head)
	if err != nil {
		return nil, err
	}

	return NewBuilder().Create().
		WithGenesis(genesis).
		WithRoot(root).
		WithHead(head).
		WithTotal(ins.Total).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createChain(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	root mined_block.Block,
	head mined_link.Link,
	total uint,
) Chain {
	out := chain{
		immutable: immutable,
		genesis:   genesis,
		root:      root,
		head:      head,
		total:     total,
	}

	return &out
}

// Hash returns the hash
func (obj *chain) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Genesis returns the genesis
func (obj *chain) Genesis() genesis.Genesis {
	return obj.genesis
}

// Root returns the root
func (obj *chain) Root() mined_block.Block {
	return obj.root
}

// Head returns the head
func (obj *chain) Head() mined_link.Link {
	return obj.head
}

// Height returns the height
func (obj *chain) Height() uint {
	return obj.Head().Link().Index()
}

// Total returns the total amount of transactions the chain contains
func (obj *chain) Total() uint {
	return obj.total
}

// CreatedOn returns the creation time
func (obj *chain) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *chain) MarshalJSON() ([]byte, error) {
	ins := createJSONChainFromChain(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *chain) UnmarshalJSON(data []byte) error {
	ins := new(JSONChain)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createChainFromJSON(ins)
	if err != nil {
		return err
	}

	insChain := pr.(*chain)
	obj.immutable = insChain.immutable
	obj.genesis = insChain.genesis
	obj.root = insChain.root
	obj.head = insChain.head
	obj.total = insChain.total
	return nil
}
