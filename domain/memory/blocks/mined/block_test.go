package mined

import (
	"encoding/json"
	"testing"
)

func TestBlock_Success(t *testing.T) {
	// mined block:
	genesisIns, minedBlockIns := CreateBlockForTestsWithoutParams()

	// repository and service:
	genService, repository, service := CreateRepositoryServiceForTests()
	err := service.Save(minedBlockIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the genesis;
	err = genService.Save(genesisIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retMinedBlock, err := repository.Retrieve(minedBlockIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !minedBlockIns.Hash().Compare(retMinedBlock.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(minedBlockIns)
	if err != nil {
		t.Errorf("the Block instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(block)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Block instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, retMinedBlock, minedBlockIns)
}
