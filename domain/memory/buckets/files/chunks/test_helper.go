package chunks

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	transfer_chunk "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CreateChunkForTests creates a chunk for tests
func CreateChunkForTests(sizeInBytes uint, data hash.Hash) Chunk {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithSizeInBytes(sizeInBytes).WithData(data).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_chunk.NewRepository(fileRepositoryService)
	trService := transfer_chunk.NewService(fileRepositoryService)
	repository := NewRepository(trRepository)
	service := NewService(repository, trService)
	return repository, service
}

// TestCompare compare two chunk instances
func TestCompare(t *testing.T, first Chunk, second Chunk) {
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
