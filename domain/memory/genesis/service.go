package genesis

import (
	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
)

type service struct {
	adapter    Adapter
	repository Repository
	trService  transfer_genesis.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	trService transfer_genesis.Service,
) Service {
	out := service{
		adapter:    adapter,
		repository: repository,
		trService:  trService,
	}

	return &out
}

// Save saves a genesis instance
func (app *service) Save(genesis Genesis) error {
	_, err := app.repository.Retrieve()
	if err == nil {
		return nil
	}

	trGenesis, err := app.adapter.ToTransfer(genesis)
	if err != nil {
		return err
	}

	return app.trService.Save(trGenesis)
}
