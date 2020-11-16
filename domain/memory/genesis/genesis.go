package genesis

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type genesis struct {
	immutable  entities.Immutable
	difficulty Difficulty
}

func createGenesisFromJSON(ins *JSONGenesis) (Genesis, error) {
	return NewBuilder().
		Create().
		WithBlockDifficultyBase(ins.BlockDifficultyBase).
		WithBlockDifficultyIncreasePerBucket(ins.BlockDifficultyIncreasePerBucket).
		WithLinkDifficulty(ins.LinkDifficulty).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createGenesis(
	immutable entities.Immutable,
	difficulty Difficulty,
) Genesis {
	out := genesis{
		immutable:  immutable,
		difficulty: difficulty,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Difficulty returns the difficulty
func (obj *genesis) Difficulty() Difficulty {
	return obj.difficulty
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
	ins := new(JSONGenesis)
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
	obj.difficulty = insGenesis.difficulty
	return nil
}
