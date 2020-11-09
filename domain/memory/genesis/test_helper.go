package genesis

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	transfer_genesis "github.com/xmn-services/buckets-network/domain/transfers/genesis"
	"github.com/xmn-services/buckets-network/libs/file"
)

// CreateGenesisForTests creates a genesis instance for tests
func CreateGenesisForTests(blockDiffBase uint, blockDiffIncreasePerTrx float64, linkDiff uint) Genesis {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().
		WithBlockDifficultyBase(blockDiffBase).
		WithBlockDifficultyIncreasePerTrx(blockDiffIncreasePerTrx).
		WithLinkDifficulty(linkDiff).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileNameWithExt := "genesis.json"
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_genesis.NewRepository(fileRepositoryService, fileNameWithExt)
	trService := transfer_genesis.NewService(fileRepositoryService, fileNameWithExt)
	repository := NewRepository(trRepository)
	service := NewService(repository, trService)
	return repository, service
}

// TestCompare compare two expense instances
func TestCompare(t *testing.T, first Genesis, second Genesis) {
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
