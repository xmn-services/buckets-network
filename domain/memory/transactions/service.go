package transactions

import (
	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	transfer_transaction "github.com/xmn-services/buckets-network/domain/transfers/transactions"
)

type service struct {
	adapter        Adapter
	repository     Repository
	addressService addresses.Service
	trService      transfer_transaction.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	addressService addresses.Service,
	trService transfer_transaction.Service,
) Service {
	out := service{
		adapter:        adapter,
		repository:     repository,
		addressService: addressService,
		trService:      trService,
	}

	return &out
}

// Save saves a transaction
func (app *service) Save(trx Transaction) error {
	hash := trx.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	if trx.HasAddress() {
		address := trx.Address()
		err = app.addressService.Save(address)
		if err != nil {
			return err
		}
	}

	trTrx, err := app.adapter.ToTransfer(trx)
	if err != nil {
		return err
	}

	return app.trService.Save(trTrx)
}

// SaveAll saves all transactions
func (app *service) SaveAll(trx []Transaction) error {
	for _, oneTrx := range trx {
		err := app.Save(oneTrx)
		if err != nil {
			return err
		}
	}

	return nil
}
