package links

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CreateLinkForTests creates a link instance for tests
func CreateLinkForTests(prev hash.Hash, next mined_blocks.Block, index uint) Link {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithPrevious(prev).WithNext(next).WithIndex(index).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (genesis.Service, Repository, Service) {
	genService, blockRepository, blockService := mined_blocks.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_link.NewRepository(fileRepositoryService)
	trService := transfer_link.NewService(fileRepositoryService)
	repository := NewRepository(blockRepository, trRepository)
	service := NewService(repository, blockService, trService)
	return genService, repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Link, second Link) {
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
