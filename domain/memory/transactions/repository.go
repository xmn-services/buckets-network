package transactions

import (
	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	transfer_transaction "github.com/xmn-services/buckets-network/domain/transfers/transactions"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	builder           Builder
	addressRepository addresses.Repository
	trRepository      transfer_transaction.Repository
}

func createRepository(
	builder Builder,
	addressRepository addresses.Repository,
	trRepository transfer_transaction.Repository,
) Repository {
	out := repository{
		builder:           builder,
		addressRepository: addressRepository,
		trRepository:      trRepository,
	}

	return &out
}

// Retrieve retrieves a transaction by hash
func (app *repository) Retrieve(hsh hash.Hash) (Transaction, error) {
	trTrx, err := app.trRepository.Retrieve(hsh)
	if err != nil {
		return nil, err
	}

	bucketHash := trTrx.Bucket()
	createdOn := trTrx.CreatedOn()
	builder := app.builder.Create().WithBucket(bucketHash).CreatedOn(createdOn)
	if trTrx.HasAddress() {
		addressHash := trTrx.Address()
		address, err := app.addressRepository.Retrieve(*addressHash)
		if err != nil {
			return nil, err
		}

		builder.WithAddress(address)
	}

	return builder.Now()
}

// RetrieveAll retrieves all trx from hashes
func (app *repository) RetrieveAll(hashes []hash.Hash) ([]Transaction, error) {
	out := []Transaction{}
	for _, oneHash := range hashes {
		trx, err := app.Retrieve(oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, trx)
	}

	return out, nil
}
