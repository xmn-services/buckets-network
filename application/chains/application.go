package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

type application struct {
	chainRepository chains.Repository
	chainService    chains.Service
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

// Update saves an updated chain
func (app *application) Update(updated chains.Chain) error {
	// retrieve the application chain:
	application, err := app.chainRepository.Retrieve()
	if err != nil {
		return err
	}

	return app.chainService.Update(application, updated)
}
