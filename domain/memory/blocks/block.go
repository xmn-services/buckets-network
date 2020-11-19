package blocks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type block struct {
	immutable  entities.Immutable
	genesis    genesis.Genesis
	additional uint
	buckets    []buckets.Bucket
}

func createBlockFromJSON(ins *JSONBlock) (Block, error) {
	genAdapter := genesis.NewAdapter()
	gen, err := genAdapter.ToGenesis(ins.Genesis)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().
		Create().
		WithGenesis(gen).
		WithAdditional(ins.Additional).
		CreatedOn(ins.CreatedOn)

	if len(ins.Buckets) > 0 {
		bucketAdapter := buckets.NewAdapter()
		buckets := []buckets.Bucket{}
		for _, oneJSBucket := range ins.Buckets {
			oneBucket, err := bucketAdapter.ToBucket(oneJSBucket)
			if err != nil {
				return nil, err
			}

			buckets = append(buckets, oneBucket)
		}

		builder.WithBuckets(buckets)
	}

	return builder.Now()
}

func createBlock(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	additional uint,
) Block {
	return createBlockInternally(immutable, genesis, additional, nil)
}

func createBlockWithBuckets(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	additional uint,
	buckets []buckets.Bucket,
) Block {
	return createBlockInternally(immutable, genesis, additional, buckets)
}

func createBlockInternally(
	immutable entities.Immutable,
	genesis genesis.Genesis,
	additional uint,
	buckets []buckets.Bucket,
) Block {
	out := block{
		immutable:  immutable,
		genesis:    genesis,
		additional: additional,
		buckets:    buckets,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Genesis returns the genesis
func (obj *block) Genesis() genesis.Genesis {
	return obj.genesis
}

// Additional returns the additional trx
func (obj *block) Additional() uint {
	return obj.additional
}

// CreatedOn returns the creation time
func (obj *block) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasBuckets returns true if there is buckets, false otherwise
func (obj *block) HasBuckets() bool {
	return obj.buckets != nil
}

// Buckets returns the buckets
func (obj *block) Buckets() []buckets.Bucket {
	return obj.buckets
}

// MarshalJSON converts the instance to JSON
func (obj *block) MarshalJSON() ([]byte, error) {
	ins := createJSONBlockFromBlock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *block) UnmarshalJSON(data []byte) error {
	ins := new(JSONBlock)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBlockFromJSON(ins)
	if err != nil {
		return err
	}

	insBlock := pr.(*block)
	obj.immutable = insBlock.immutable
	obj.genesis = insBlock.genesis
	obj.additional = insBlock.additional
	obj.buckets = insBlock.buckets
	return nil
}
