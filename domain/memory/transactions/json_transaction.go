package transactions

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
)

// JSONTransaction represents a json transaction
type JSONTransaction struct {
	Address   *addresses.JSONAddress `json:"address"`
	Bucket    string                 `json:"bucket"`
	CreatedOn time.Time              `json:"created_on"`
}

func createJSONTransactionFromTransaction(transaction Transaction) *JSONTransaction {
	var address *addresses.JSONAddress
	if transaction.HasAddress() {
		addressAdapter := addresses.NewAdapter()
		addr := transaction.Address()
		address = addressAdapter.ToJSON(addr)
	}

	bucket := transaction.Bucket().String()
	createdOn := transaction.CreatedOn()
	return createJSONTransaction(address, bucket, createdOn)
}

func createJSONTransaction(
	address *addresses.JSONAddress,
	bucket string,
	createdOn time.Time,
) *JSONTransaction {
	out := JSONTransaction{
		Address:   address,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
