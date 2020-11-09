package genesis

import (
	"time"
)

// JSONGenesis represents a JSON genesis
type JSONGenesis struct {
	LinkDifficulty                uint      `json:"link_difficulty"`
	BlockDifficultyBase           uint      `json:"block_difficulty_base"`
	BlockDifficultyIncreasePerTrx float64   `json:"block_difficulty_increase_per_transaction"`
	CreatedOn                     time.Time `json:"created_on"`
}

func createJSONGenesisFromGenesis(gen Genesis) *JSONGenesis {
	difficulty := gen.Difficulty()
	linkDifficulty := difficulty.Link()
	blockDifficulty := difficulty.Block()
	blockDifficultyBase := blockDifficulty.Base()
	blockDifficultyIncreasePerTrx := blockDifficulty.IncreasePerTrx()
	createdOn := gen.CreatedOn()
	return createJSONGenesis(
		linkDifficulty,
		blockDifficultyBase,
		blockDifficultyIncreasePerTrx,
		createdOn,
	)
}

func createJSONGenesis(
	linkDifficulty uint,
	blockDifficultyBase uint,
	blockDifficultyIncreasePerTrx float64,
	createdOn time.Time,
) *JSONGenesis {
	out := JSONGenesis{
		LinkDifficulty:                linkDifficulty,
		BlockDifficultyBase:           blockDifficultyBase,
		BlockDifficultyIncreasePerTrx: blockDifficultyIncreasePerTrx,
		CreatedOn:                     createdOn,
	}

	return &out
}
