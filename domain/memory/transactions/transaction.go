package transactions

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type transaction struct {
	immutable entities.Immutable
	bucket    hash.Hash
	address   addresses.Address
}

func createTransactionFromJSON(js *JSONTransaction) (Transaction, error) {
	hashAdapter := hash.NewAdapter()
	bucket, err := hashAdapter.FromString(js.Bucket)
	if err != nil {
		return nil, err
	}

	builder := NewBuilder().Create().WithBucket(*bucket).CreatedOn(js.CreatedOn)
	if js.Address != nil {
		addressAdapter := addresses.NewAdapter()
		address, err := addressAdapter.ToAddress(js.Address)
		if err != nil {
			return nil, err
		}

		builder.WithAddress(address)
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
	address addresses.Address,
) Transaction {
	return createTransactionInternally(immutable, bucket, address)
}

func createTransactionInternally(
	immutable entities.Immutable,
	bucket hash.Hash,
	address addresses.Address,
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

// Bucket returns the bucket
func (obj *transaction) Bucket() hash.Hash {
	return obj.bucket
}

// HasAddress returns true if there is an address, false otherwise
func (obj *transaction) HasAddress() bool {
	return obj.address != nil
}

// Address returns the address, if any
func (obj *transaction) Address() addresses.Address {
	return obj.address
}

// CreatedOn returns the creation time
func (obj *transaction) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *transaction) MarshalJSON() ([]byte, error) {
	ins := createJSONTransactionFromTransaction(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *transaction) UnmarshalJSON(data []byte) error {
	ins := new(JSONTransaction)
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
