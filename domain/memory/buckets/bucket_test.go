package buckets

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestFile_Success(t *testing.T) {

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromBytes([]byte("this is the first hash"))
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

	bucketIns := CreateBucketForTests(files)

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(bucketIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the file:
	err = service.Save(bucketIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBucket, err := repository.Retrieve(bucketIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bucketIns.Hash().Compare(retBucket.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(bucketIns)
	if err != nil {
		t.Errorf("the File instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(bucket)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a File instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, bucketIns, retBucket)
}
