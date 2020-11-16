package genesis

import (
	"github.com/xmn-services/buckets-network/application/miners"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

type application struct {
	minerApplication miners.Application
	chainService     chains.Service
}

func createApplication(
	minerApplication miners.Application,
	chainService chains.Service,
) Application {
	out := application{
		minerApplication: minerApplication,
		chainService:     chainService,
	}

	return &out
}

// Init initializes the genesis block
func (app *application) Init(baseDifficulty uint, increasePerBucket float64, linkDifficulty uint) error {
	// mine the chain:
	chain, err := app.minerApplication.Init(baseDifficulty, increasePerBucket, linkDifficulty)
	if err != nil {
		return err
	}

	// save the chain:
	return app.chainService.Insert(chain)
}
