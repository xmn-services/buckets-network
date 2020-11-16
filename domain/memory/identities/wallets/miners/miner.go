package miners

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type miner struct {
	mutable     entities.Mutable
	toTransact  buckets.Buckets
	queue       buckets.Buckets
	broadcasted permanents.Buckets
	toLink      blocks.Blocks
}

func createMinerFromJSON(ins *JSONMiner) (Miner, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	bucketsAdapter := buckets.NewAdapter()
	toTransact, err := bucketsAdapter.ToBuckets(ins.ToTransact)
	if err != nil {
		return nil, err
	}

	queue, err := bucketsAdapter.ToBuckets(ins.Queue)
	if err != nil {
		return nil, err
	}

	permanentBucketsAdapter := permanents.NewAdapter()
	broadcasted, err := permanentBucketsAdapter.ToBuckets(ins.Broadcasted)
	if err != nil {
		return nil, err
	}

	blockAdapter := blocks.NewAdapter()
	toLink, err := blockAdapter.ToBlocks(ins.ToLink)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithToTransact(toTransact).
		WithQueue(queue).
		WithBroadcasted(broadcasted).
		WithToLink(toLink).
		CreatedOn(ins.CreatedOn).
		LastUpdatedOn(ins.LastUpdatedOn).
		Now()
}

func createMiner(
	mutable entities.Mutable,
	toTransact buckets.Buckets,
	queue buckets.Buckets,
	broadcasted permanents.Buckets,
	toLink blocks.Blocks,
) Miner {
	out := miner{
		mutable:     mutable,
		toTransact:  toTransact,
		queue:       queue,
		broadcasted: broadcasted,
		toLink:      toLink,
	}

	return &out
}

// Hash returns the hash
func (obj *miner) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// ToTransact returns the toTransact buckets
func (obj *miner) ToTransact() buckets.Buckets {
	return obj.toTransact
}

// Queue returns the queue buckets
func (obj *miner) Queue() buckets.Buckets {
	return obj.queue
}

// Broadcasted returns the broadcasted buckets
func (obj *miner) Broadcasted() permanents.Buckets {
	return obj.broadcasted
}

// ToLink returns the toLink blocks
func (obj *miner) ToLink() blocks.Blocks {
	return obj.toLink
}

// CreatedOn returns the creation time
func (obj *miner) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *miner) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *miner) MarshalJSON() ([]byte, error) {
	ins := createJSONMinerFromMiner(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *miner) UnmarshalJSON(data []byte) error {
	ins := new(JSONMiner)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createMinerFromJSON(ins)
	if err != nil {
		return err
	}

	insMiner := pr.(*miner)
	obj.mutable = insMiner.mutable
	obj.toTransact = insMiner.toTransact
	obj.queue = insMiner.queue
	obj.broadcasted = insMiner.broadcasted
	obj.toLink = insMiner.toLink
	return nil
}
