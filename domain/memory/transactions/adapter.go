package transactions

import (
	transfer_transaction "github.com/xmn-services/buckets-network/domain/transfers/transactions"
)

type adapter struct {
	trBuilder transfer_transaction.Builder
}

func createAdapter(
	trBuilder transfer_transaction.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a transaction to a transfer transaction instance
func (app *adapter) ToTransfer(trx Transaction) (transfer_transaction.Transaction, error) {
	hsh := trx.Hash()
	bucket := trx.Bucket()
	createdOn := trx.CreatedOn()

	builder := app.trBuilder.Create().
		WithHash(hsh).
		WithBucket(bucket).
		CreatedOn(createdOn)

	if trx.HasAddress() {
		addressHash := trx.Address().Hash()
		builder.WithAddress(addressHash)
	}

	return builder.Now()
}

// ToJSON converts a transaction to a JSON instances
func (app *adapter) ToJSON(trx Transaction) *JSONTransaction {
	return createJSONTransactionFromTransaction(trx)
}

// ToTransaction converts a JSON transaction to a Transaction instances
func (app *adapter) ToTransaction(ins *JSONTransaction) (Transaction, error) {
	return createTransactionFromJSON(ins)
}
