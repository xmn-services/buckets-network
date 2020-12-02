package buckets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	libs_file "github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
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

// CreateBucketForTestsWithoutParams creates a bucket for tests without params
func CreateBucketForTestsWithoutParams() Bucket {
	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromBytes([]byte(fmt.Sprintf("this is the first hash: %s", time.Now().UTC().String())))
	secondHash, _ := hashAdapter.FromBytes([]byte("this is the second hash"))

	firstChunks := []chunks.Chunk{
		chunks.CreateChunkForTests(uint(345234), *firstHash),
	}

	secondChunks := []chunks.Chunk{
		chunks.CreateChunkForTests(uint(2345234), *secondHash),
	}

	firstRelativePath := "/first/is/relative/path"
	secondRelativePath := "/second/is/relative/path"

	files := []files.File{
		files.CreateFileForTests(firstRelativePath, firstChunks),
		files.CreateFileForTests(secondRelativePath, secondChunks),
	}

	return CreateBucketForTests(files)
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
