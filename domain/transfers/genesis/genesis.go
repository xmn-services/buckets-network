package genesis

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type genesis struct {
	immutable               entities.Immutable
	blockDiffBase           uint
	blockDiffIncreasePerBucket float64
	linkDiff                uint
}

func createGenesisFromJSON(ins *jsonGenesis) (Genesis, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithBlockDifficultyBase(ins.BlockDiffBase).
		WithBlockDifficultyIncreasePerBucket(ins.BlockDiffIncreasePerBucket).
		WithLinkDifficulty(ins.LinkDiff).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createGenesis(
	immutable entities.Immutable,
	blockDiffBase uint,
	blockDiffIncreasePerBucket float64,
	linkDiff uint,
) Genesis {
	out := genesis{
		immutable:               immutable,
		blockDiffBase:           blockDiffBase,
		blockDiffIncreasePerBucket: blockDiffIncreasePerBucket,
		linkDiff:                linkDiff,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// BlockDifficultyBase returns the block difficulty base
func (obj *genesis) BlockDifficultyBase() uint {
	return obj.blockDiffBase
}

// BlockDifficultyIncreasePerBucket returns the block difficulty increase per bucket
func (obj *genesis) BlockDifficultyIncreasePerBucket() float64 {
	return obj.blockDiffIncreasePerBucket
}

// LinkDifficulty returns the link difficulty
func (obj *genesis) LinkDifficulty() uint {
	return obj.linkDiff
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *genesis) MarshalJSON() ([]byte, error) {
	ins := createJSONGenesisFromGenesis(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *genesis) UnmarshalJSON(data []byte) error {
	ins := new(jsonGenesis)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createGenesisFromJSON(ins)
	if err != nil {
		return err
	}

	insGenesis := pr.(*genesis)
	obj.immutable = insGenesis.immutable
	obj.blockDiffBase = insGenesis.blockDiffBase
	obj.blockDiffIncreasePerBucket = insGenesis.blockDiffIncreasePerBucket
	obj.linkDiff = insGenesis.linkDiff
	return nil
}
