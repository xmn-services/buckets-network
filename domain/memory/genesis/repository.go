package genesis

import (
	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
)

type repository struct {
	trRepository transfer_genesis.Repository
	builder      Builder
}

func createRepository(
	builder Builder,
	trRepository transfer_genesis.Repository,
) Repository {
	out := repository{
		builder:      builder,
		trRepository: trRepository,
	}

	return &out
}

// Retrieve retrieves a genesis instance
func (app *repository) Retrieve() (Genesis, error) {
	trGen, err := app.trRepository.Retrieve()
	if err != nil {
		return nil, err
	}

	blockDiffBase := trGen.BlockDifficultyBase()
	blockDiffIncreasePerTrx := trGen.BlockDifficultyIncreasePerTrx()
	linkDiff := trGen.LinkDifficulty()
	createdOn := trGen.CreatedOn()
	return app.builder.Create().
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx).
		WithLinkDifficulty(linkDiff).
		CreatedOn(createdOn).
		Now()
}
