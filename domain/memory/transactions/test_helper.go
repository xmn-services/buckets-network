package transactions

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	transfer_transaction "github.com/xmn-services/buckets-network/domain/transfers/transactions"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CreateTransactionForTests creates a transaction for tests
func CreateTransactionForTests(amountPubKeyInRing uint, executesOnTrigger bool, bucket hash.Hash) Transaction {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().
		Create().
		WithBucket(bucket).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests(amountPubKeyInRing uint) (Repository, Service) {
	addressRepository, addressService := addresses.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_transaction.NewRepository(fileRepositoryService)
	trService := transfer_transaction.NewService(fileRepositoryService)
	repository := NewRepository(addressRepository, trRepository)
	service := NewService(repository, addressService, trService)
	return repository, service
}

// TestCompare compare two instances
func TestCompare(t *testing.T, first Transaction, second Transaction) {
	js, err := json.Marshal(first)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = json.Unmarshal(js, second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reJS, err := json.Marshal(second)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(js, reJS) != 0 {
		t.Errorf("the transformed javascript is different.\n%s\n%s", js, reJS)
		return
	}

	if !first.Hash().Compare(second.Hash()) {
		t.Errorf("the instance conversion failed")
		return
	}
}
