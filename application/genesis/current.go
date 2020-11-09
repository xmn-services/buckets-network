package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/genesis"
)

type current struct {
	genesisBuilder    genesis.Builder
	genesisRepository genesis.Repository
	genesisService    genesis.Service
}

func createCurrent(
	genesisBuilder genesis.Builder,
	genesisRepository genesis.Repository,
	genesisService genesis.Service,
) Current {
	out := current{
		genesisBuilder:    genesisBuilder,
		genesisRepository: genesisRepository,
		genesisService:    genesisService,
	}

	return &out
}

// Init initializes the genesis block
func (app *current) Init(
	blockDifficultyBase uint,
	blockDifficultyIncreasePerTrx float64,
	linkDifficulty uint,
) error {
	_, err := app.genesisRepository.Retrieve()
	if err == nil {
		return errors.New("the genesis block has already been created")
	}

	createdOn := time.Now().UTC()
	gen, err := app.genesisBuilder.Create().
		WithBlockDifficultyBase(blockDifficultyBase).
		WithBlockDifficultyIncreasePerTrx(blockDifficultyIncreasePerTrx).
		WithLinkDifficulty(linkDifficulty).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	return app.genesisService.Save(gen)
}
