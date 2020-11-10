package buckets

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	libs_file "github.com/xmn-services/buckets-network/libs/file"
)

// CreateBucketForTests creates a bucket for tests
func CreateBucketForTests(files []files.File) Bucket {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithFiles(files).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileRepositoryService := libs_file.CreateRepositoryServiceForTests()
	fileRepository, fileService := files.CreateRepositoryServiceForTests()
	trRepository := transfer_bucket.NewRepository(fileRepositoryService)
	trService := transfer_bucket.NewService(fileRepositoryService)
	repository := NewRepository(fileRepository, trRepository)
	service := NewService(fileService, repository, trService)
	return repository, service
}

// TestCompare compare two bucket instances
func TestCompare(t *testing.T, first Bucket, second Bucket) {
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
