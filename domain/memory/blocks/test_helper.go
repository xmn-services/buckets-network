package blocks

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
	"github.com/xmn-services/buckets-network/libs/file"
)

// CreateBlockForTests creates a block for tests
func CreateBlockForTests(gen genesis.Genesis, additional uint, buckets []buckets.Bucket) Block {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().
		WithGenesis(gen).
		WithAdditional(additional).
		WithBuckets(buckets).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (genesis.Service, Repository, Service) {
	genesisRepository, genesisService := genesis.CreateRepositoryServiceForTests()
	bucketRepository, bucketService := buckets.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_block.NewRepository(fileRepositoryService)
	trService := transfer_block.NewService(fileRepositoryService)
	repository := NewRepository(genesisRepository, bucketRepository, trRepository)
	service := NewService(repository, bucketService, trService)
	return genesisService, repository, service
}

// TestCompare compare two block instances
func TestCompare(t *testing.T, first Block, second Block) {
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
