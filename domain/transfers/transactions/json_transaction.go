package transactions

import (
	"time"
)

type jsonTransaction struct {
	Hash      string    `json:"hash"`
	Address   string    `json:"address"`
	Bucket    string    `json:"bucket"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONTransactionFromTransaction(ins Transaction) *jsonTransaction {
	hash := ins.Hash().String()
	bucket := ins.Bucket().String()

	address := ""
	if ins.HasAddress() {
		address = ins.Address().String()
	}

	createdOn := ins.CreatedOn()
	return createJSONTransaction(hash, address, bucket, createdOn)
}

func createJSONTransaction(
	hash string,
	address string,
	bucket string,
	createdOn time.Time,
) *jsonTransaction {
	out := jsonTransaction{
		Hash:      hash,
		Address:   address,
		Bucket:    bucket,
		CreatedOn: createdOn,
	}

	return &out
}
