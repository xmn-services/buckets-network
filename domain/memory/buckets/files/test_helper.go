package files

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	transfer_file "github.com/xmn-services/buckets-network/domain/transfers/buckets/files"
	libs_file "github.com/xmn-services/buckets-network/libs/file"
)

// CreateFileForTests creates a file for tests
func CreateFileForTests(relativePath string, chunks []chunks.Chunk) File {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithRelativePath(relativePath).WithChunks(chunks).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileRepositoryService := libs_file.CreateRepositoryServiceForTests()
	chunkRepository, chunkService := chunks.CreateRepositoryServiceForTests()
	trRepository := transfer_file.NewRepository(fileRepositoryService)
	trService := transfer_file.NewService(fileRepositoryService)
	repository := NewRepository(chunkRepository, trRepository)
	service := NewService(chunkService, repository, trService)
	return repository, service
}

// TestCompare compare two file instances
func TestCompare(t *testing.T, first File, second File) {
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
