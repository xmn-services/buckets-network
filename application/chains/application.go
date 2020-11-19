package chains

import (
	"github.com/xmn-services/buckets-network/application/miners"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

type application struct {
	minerApplication miners.Application
	chainRepository  chains.Repository
	chainService     chains.Service
	chainBuilder     chains.Builder
}

func createApplication(
	minerApplication miners.Application,
	chainRepository chains.Repository,
	chainService chains.Service,
	chainBuilder chains.Builder,
) Application {
	out := application{
		minerApplication: minerApplication,
		chainRepository:  chainRepository,
		chainService:     chainService,
		chainBuilder:     chainBuilder,
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

// Update updates the head of the chain:
func (app *application) Update(newMinedLink mined_link.Link) error {
	// retrieve the application chain:
	original, err := app.chainRepository.Retrieve()
	if err != nil {
		return err
	}

	gen := original.Genesis()
	root := original.Root()
	total := original.Total() + 1
	updatedChain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(newMinedLink).WithTotal(total).Now()
	if err != nil {
		return err
	}

	return app.chainService.Update(original, updatedChain)
}
