package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	bucket    hash.Hash
	address   *hash.Hash
}

func createTransactionFromJSON(ins *jsonTransaction) (Transaction, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	bucket, err := hashAdapter.FromString(ins.Bucket)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().
		Create().
		WithHash(*hsh).
		WithBucket(*bucket).
		CreatedOn(ins.CreatedOn)

	if ins.Address != "" {
		address, err := hashAdapter.FromString(ins.Address)
		if err != nil {
			return nil, err
		}

		builder.WithAddress(*address)
	}

	return builder.Now()
}

func createTransaction(
	immutable entities.Immutable,
	bucket hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, bucket, nil)
}

func createTransactionWithAddress(
	immutable entities.Immutable,
	bucket hash.Hash,
	address *hash.Hash,
) Transaction {
	return createTransactionInternally(immutable, bucket, address)
}

func createTransactionInternally(
	immutable entities.Immutable,
	bucket hash.Hash,
	address *hash.Hash,
) Transaction {
	out := transaction{
		immutable: immutable,
		bucket:    bucket,
		address:   address,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bucket returns the bucket hash
func (obj *transaction) Bucket() hash.Hash {
	return obj.bucket
}

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasAddress returns true if the transaction contains an address, false otherwise
func (obj *transaction) HasAddress() bool {
	return obj.address != nil
}

// Address returns the address hash, if any
func (obj *transaction) Address() *hash.Hash {
	return obj.address
}

// MarshalJSON converts the instance to JSON
func (obj *transaction) MarshalJSON() ([]byte, error) {
	ins := createJSONTransactionFromTransaction(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *transaction) UnmarshalJSON(data []byte) error {
	ins := new(jsonTransaction)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createTransactionFromJSON(ins)
	if err != nil {
		return err
	}

	insTransaction := pr.(*transaction)
	obj.immutable = insTransaction.immutable
	obj.address = insTransaction.address
	obj.bucket = insTransaction.bucket
	return nil
}
