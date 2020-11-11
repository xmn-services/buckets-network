package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/genesis"
)

type application struct {
	genesisBuilder    genesis.Builder
	genesisRepository genesis.Repository
	genesisService    genesis.Service
}

func createApplication(
	genesisBuilder genesis.Builder,
	genesisRepository genesis.Repository,
	genesisService genesis.Service,
) Application {
	out := application{
		genesisBuilder:    genesisBuilder,
		genesisRepository: genesisRepository,
		genesisService:    genesisService,
	}

	return &out
}

// Init initializes the genesis block
func (app *application) Init(
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
