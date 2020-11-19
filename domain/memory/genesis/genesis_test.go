package genesis

import (
	"encoding/json"
	"testing"
)

func TestGenesis_Success(t *testing.T) {
	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerBucket := float64(1.0)
	linkDiff := uint(1)
	miningValue := uint8(6)
	genesisIns := CreateGenesisForTests(miningValue, blockDiffBase, blockDiffIncreasePerBucket, linkDiff)

	// repository and service:
	repository, service := CreateRepositoryServiceForTests()
	err := service.Save(genesisIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retGenesis, err := repository.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !genesisIns.Hash().Compare(retGenesis.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(genesisIns)
	if err != nil {
		t.Errorf("the Genesis instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(genesis)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Genesis instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, genesisIns)
}
