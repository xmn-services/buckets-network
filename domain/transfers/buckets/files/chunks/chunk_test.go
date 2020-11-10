package chunks

import (
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestChunk_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	sizeInBytes := uint(56353)
	data, _ := hashAdapter.FromBytes([]byte("this is some data"))
	createdOn := time.Now().UTC()

	chk, err := NewBuilder().Create().WithHash(*hsh).WithSizeInBytes(sizeInBytes).WithData(*data).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !chk.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !chk.Data().Compare(*data) {
		t.Errorf("the data hash is invalid")
		return
	}

	if chk.SizeInBytes() != sizeInBytes {
		t.Errorf("the sizeInBytes was expected to be %d, %d returned", sizeInBytes, chk.SizeInBytes())
		return
	}

	if !chk.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), chk.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

	err = service.Save(chk)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retChunk, err := repository.Retrieve(chk.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, chk, retChunk)
}
