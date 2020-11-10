package chunks

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestChunk_Success(t *testing.T) {
	sizeInBytes := uint(2345)
	data, _ := hash.NewAdapter().FromBytes([]byte("this is some data"))

	chk := CreateChunkForTests(sizeInBytes, *data)

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(chk)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the chunk:
	err = service.Save(chk)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChk, err := repository.Retrieve(chk.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !chk.Hash().Compare(retChk.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(chk)
	if err != nil {
		t.Errorf("the Chunk instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(chunk)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Chunk instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, chk, retChk)
}
