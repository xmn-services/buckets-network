package transactions

import (
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestTransaction_hasBucket_hasAddress_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	bucket, _ := hashAdapter.FromBytes([]byte("to build the bucket hash..."))
	address, _ := hashAdapter.FromBytes([]byte("to build the address hash..."))

	createdOn := time.Now().UTC()
	transaction, err := NewBuilder().
		Create().
		WithHash(*hsh).
		WithBucket(*bucket).
		WithAddress(*address).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !transaction.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !transaction.Bucket().Compare(*bucket) {
		t.Errorf("the bucket is invalid")
		return
	}

	if !transaction.HasAddress() {
		t.Errorf("the address was expected to be valid")
		return
	}

	if !transaction.Address().Compare(*address) {
		t.Errorf("the address is invalid")
		return
	}

	if !transaction.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), transaction.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(transaction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTransaction, err := repository.Retrieve(transaction.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, transaction, retTransaction)
}
