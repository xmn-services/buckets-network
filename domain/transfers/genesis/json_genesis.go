package genesis

import (
	"time"
)

type jsonGenesis struct {
	Hash                       string    `json:"hash"`
	MiningValue                uint8     `json:"mining_value"`
	BlockDiffBase              uint      `json:"block_difficulty_base"`
	BlockDiffIncreasePerBucket float64   `json:"block_difficulty_increase_per_bucket"`
	LinkDiff                   uint      `json:"link_difficulty"`
	CreatedOn                  time.Time `json:"created_on"`
}

func createJSONGenesisFromGenesis(ins Genesis) *jsonGenesis {
	hash := ins.Hash().String()
	miningValue := ins.MiningValue()
	blockDiffBase := ins.BlockDifficultyBase()
	blockDiffIncreasePerBucket := ins.BlockDifficultyIncreasePerBucket()
	linkDiff := ins.LinkDifficulty()
	createdOn := ins.CreatedOn()
	return createJSONGenesis(hash, miningValue, blockDiffBase, blockDiffIncreasePerBucket, linkDiff, createdOn)
}

func createJSONGenesis(
	hash string,
	miningValue uint8,
	blockDiffBase uint,
	blockDiffIncreasePerBucket float64,
	linkDiff uint,
	createdOn time.Time,
) *jsonGenesis {
	out := jsonGenesis{
		Hash:                       hash,
		MiningValue:                miningValue,
		BlockDiffBase:              blockDiffBase,
		BlockDiffIncreasePerBucket: blockDiffIncreasePerBucket,
		LinkDiff:                   linkDiff,
		CreatedOn:                  createdOn,
	}

	return &out
}
