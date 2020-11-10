package buckets

import (
	"fmt"
	"testing"
	"time"

	lib_file "github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

func TestBucket_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))

	data := [][]byte{}
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("to build the %d file hash...", i)
		oneTrx, _ := hashAdapter.FromBytes([]byte(str))
		data = append(data, oneTrx.Bytes())
	}

	files, _ := hashtree.NewBuilder().Create().WithBlocks(data).Now()
	amount := uint(len(data))
	createdOn := time.Now().UTC()

	bucketIns, err := NewBuilder().Create().WithHash(*hsh).WithFiles(files).WithAmount(amount).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bucketIns.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if !bucketIns.Files().Head().Compare(files.Head()) {
		t.Errorf("the hashtree is invalid")
		return
	}

	if bucketIns.Amount() != amount {
		t.Errorf("the amount was expected to be %d, %d returned", amount, bucketIns.Amount())
		return
	}

	if !bucketIns.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), bucketIns.CreatedOn().String())
		return
	}

	// repository and service:
	fileRepositoryService := lib_file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)

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

	// compare:
	TestCompare(t, bucketIns, retBucket)
}
