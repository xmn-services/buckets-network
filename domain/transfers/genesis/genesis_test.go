package genesis

import (
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestGenesis_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	hsh, _ := hashAdapter.FromBytes([]byte("to build the hash..."))
	blockDiffBase := uint(45)
	blockDiffIncreasePerBucket := float64(0.0021)
	linkDiff := uint(2)
	createdOn := time.Now().UTC()

	genesis, err := NewBuilder().Create().WithHash(*hsh).WithBlockDifficultyBase(blockDiffBase).WithBlockDifficultyIncreasePerBucket(blockDiffIncreasePerBucket).WithLinkDifficulty(linkDiff).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !genesis.Hash().Compare(*hsh) {
		t.Errorf("the hash is invalid")
		return
	}

	if genesis.BlockDifficultyBase() != blockDiffBase {
		t.Errorf("the blockDifficultyBase is invalid, expected: %d, returned: %d", blockDiffBase, genesis.BlockDifficultyBase())
		return
	}

	if genesis.BlockDifficultyIncreasePerBucket() != blockDiffIncreasePerBucket {
		t.Errorf("the blockDifficultyIncreasePerBucket is invalid, expected: %f, returned: %f", blockDiffIncreasePerBucket, genesis.BlockDifficultyIncreasePerBucket())
		return
	}

	if genesis.LinkDifficulty() != linkDiff {
		t.Errorf("the linkDifficulty is invalid, expected: %d, returned: %d", linkDiff, genesis.LinkDifficulty())
		return
	}

	if !genesis.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time is invalid, expected: %s, returned: %s", createdOn.String(), genesis.CreatedOn().String())
		return
	}

	// repository and service:
	name := "genesis.json"
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService, name)
	service := NewService(fileRepositoryService, name)

	err = service.Save(genesis)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retGenesis, err := repository.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, genesis, retGenesis)
}
