package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

type application struct {
	chainRepository chains.Repository
	chainService    chains.Service
	chainBuilder    chains.Builder
}

func createApplication(
	chainRepository chains.Repository,
	chainService chains.Service,
) Application {
	out := application{
		chainRepository: chainRepository,
		chainService:    chainService,
	}

	return &out
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
