package files

import (
	"fmt"
	"testing"
	"time"

	lib_file "github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

func TestFile_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	relativePath := "/some/path"

	data := [][]byte{}
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("to build the %d chunk hash...", i)
		oneTrx, _ := hashAdapter.FromBytes([]byte(str))
		data = append(data, oneTrx.Bytes())
	}

	chunks, _ := hashtree.NewBuilder().Create().WithBlocks(data).Now()
	amount := uint(len(data))
	createdOn := time.Now().UTC()

	fileIns, err := NewBuilder().Create().WithHash(*hsh).WithRelativePath(relativePath).WithChunks(chunks).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !fileIns.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !fileIns.Chunks().Head().Compare(chunks.Head()) {
		t.Errorf("the hashtree is invalid")
		return
	}

	if fileIns.RelativePath() != relativePath {
		t.Errorf("the relativePath was expected to be %s, %s returned", relativePath, fileIns.RelativePath())
		return
	}

	if fileIns.Amount() != amount {
		t.Errorf("the amount was expected to be %d, %d returned", amount, fileIns.Amount())
		return
	}

	if !fileIns.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), fileIns.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := lib_file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

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

	// compare:
	TestCompare(t, fileIns, retFile)
}
