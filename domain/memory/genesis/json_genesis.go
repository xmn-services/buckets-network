package genesis

import (
	"time"
)

// JSONGenesis represents a JSON genesis
type JSONGenesis struct {
	MiningValue                      uint8     `json:"mining_value"`
	LinkDifficulty                   uint      `json:"link_difficulty"`
	BlockDifficultyBase              uint      `json:"block_difficulty_base"`
	BlockDifficultyIncreasePerBucket float64   `json:"block_difficulty_increase_per_transaction"`
	CreatedOn                        time.Time `json:"created_on"`
}

func createJSONGenesisFromGenesis(gen Genesis) *JSONGenesis {
	miningValue := gen.MiningValue()
	difficulty := gen.Difficulty()
	linkDifficulty := difficulty.Link()
	blockDifficulty := difficulty.Block()
	blockDifficultyBase := blockDifficulty.Base()
	blockDifficultyIncreasePerBucket := blockDifficulty.IncreasePerBucket()
	createdOn := gen.CreatedOn()
	return createJSONGenesis(
		miningValue,
		linkDifficulty,
		blockDifficultyBase,
		blockDifficultyIncreasePerBucket,
		createdOn,
	)
}

func createJSONGenesis(
	miningValue uint8,
	linkDifficulty uint,
	blockDifficultyBase uint,
	blockDifficultyIncreasePerBucket float64,
	createdOn time.Time,
) *JSONGenesis {
	out := JSONGenesis{
		MiningValue:                      miningValue,
		LinkDifficulty:                   linkDifficulty,
		BlockDifficultyBase:              blockDifficultyBase,
		BlockDifficultyIncreasePerBucket: blockDifficultyIncreasePerBucket,
		CreatedOn:                        createdOn,
	}

	return &out
}
