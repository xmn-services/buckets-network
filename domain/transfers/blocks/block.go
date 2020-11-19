package blocks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type block struct {
	immutable  entities.Immutable
	buckets    hashtree.HashTree
	amount     uint
	additional uint
}

func createBlockFromJSON(ins *jsonBlock) (Block, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().
		Create().
		WithHash(*hsh).
		WithAmount(ins.Amount).
		WithAdditional(ins.Additional).
		CreatedOn(ins.CreatedOn)

	if ins.Buckets != nil {
		compact, err := hashtree.NewAdapter().FromJSON(ins.Buckets)
		if err != nil {
			return nil, err
		}

		buckets, err := compact.Leaves().HashTree()
		if err != nil {
			return nil, err
		}

		builder.WithBuckets(buckets)
	}

	return builder.Now()
}

func createBlock(
	immutable entities.Immutable,
	amount uint,
	additional uint,
) Block {
	return createBlockInternally(immutable, amount, additional, nil)
}

func createBlockWithBuckets(
	immutable entities.Immutable,
	amount uint,
	additional uint,
	buckets hashtree.HashTree,
) Block {
	return createBlockInternally(immutable, amount, additional, buckets)
}

func createBlockInternally(
	immutable entities.Immutable,
	amount uint,
	additional uint,
	buckets hashtree.HashTree,
) Block {
	out := block{
		immutable:  immutable,
		amount:     amount,
		additional: additional,
		buckets:    buckets,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Amount returns the amount
func (obj *block) Amount() uint {
	return obj.amount
}

// Additional returns the additional
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

// Buckets returns the buckets hashtree
func (obj *block) Buckets() hashtree.HashTree {
	return obj.buckets
}

// MarshalJSON converts the instance to JSON
func (obj *block) MarshalJSON() ([]byte, error) {
	ins := createJSONBlockFromBlock(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *block) UnmarshalJSON(data []byte) error {
	ins := new(jsonBlock)
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
	obj.buckets = insBlock.buckets
	obj.amount = insBlock.amount
	obj.additional = insBlock.additional
	return nil
}
