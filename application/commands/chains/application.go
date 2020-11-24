package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

type application struct {
	chainRepository chains.Repository
}

func createApplication(
	chainRepository chains.Repository,
) Application {
	out := application{
		chainRepository: chainRepository,
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
