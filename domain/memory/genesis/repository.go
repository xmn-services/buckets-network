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

	miningValue := trGen.MiningValue()
	blockDiffBase := trGen.BlockDifficultyBase()
	blockDiffIncreasePerBucket := trGen.BlockDifficultyIncreasePerBucket()
	linkDiff := trGen.LinkDifficulty()
	createdOn := trGen.CreatedOn()
	return app.builder.Create().
		WithMiningValue(miningValue).
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerBucket(blockDiffIncreasePerBucket).
		WithLinkDifficulty(linkDiff).
		CreatedOn(createdOn).
		Now()
}
