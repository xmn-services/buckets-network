package files

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestFile_Success(t *testing.T) {

	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromBytes([]byte("this is the first hash"))
	secondHash, _ := hashAdapter.FromBytes([]byte("this is the second hash"))

	chunks := []chunks.Chunk{
		chunks.CreateChunkForTests(uint(345234), *firstHash),
		chunks.CreateChunkForTests(uint(2345234), *secondHash),
	}

	relativePath := "/this/is/relative/path"
	fileIns := CreateFileForTests(relativePath, chunks)

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(fileIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the file:
	err = service.Save(fileIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFile, err := repository.Retrieve(fileIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !fileIns.Hash().Compare(retFile.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(fileIns)
	if err != nil {
		t.Errorf("the File instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(file)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a File instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, fileIns, retFile)
}
