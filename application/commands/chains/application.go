package chains

import (
	"github.com/xmn-services/buckets-network/application/commands/miners"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

type application struct {
	minerApplication miners.Application
	chainRepository  chains.Repository
	chainService     chains.Service
}

func createApplication(
	minerApplication miners.Application,
	chainRepository chains.Repository,
	chainService chains.Service,
) Application {
	out := application{
		minerApplication: minerApplication,
		chainRepository:  chainRepository,
		chainService:     chainService,
	}

	return &out
}

// Init initializes the chain
func (app *application) Init(miningValue uint8, baseDifficulty uint, increasePerBucket float64, linkDifficulty uint, rootAdditionalBuckets uint, headAdditionalBuckets uint) error {
	// mine the chain:
	chain, err := app.minerApplication.Init(miningValue, baseDifficulty, increasePerBucket, linkDifficulty, rootAdditionalBuckets, headAdditionalBuckets)
	if err != nil {
		return err
	}

	// save the chain:
	return app.chainService.Insert(chain)
}

// Retrieve retrieves the chain
func (app *application) Retrieve() (chains.Chain, error) {
	return app.chainRepository.Retrieve()
}

// RetrieveAtIndex retrieves the chain at index
func (app *application) RetrieveAtIndex(index uint) (chains.Chain, error) {
	return app.chainRepository.RetrieveAtIndex(index)
}
