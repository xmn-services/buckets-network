package filedisks

import (
	"os"
	"reflect"
	"testing"
)

func TestInit_Success(t *testing.T) {
	miningValue := uint8(0)
	baseDifficulty := uint(1)
	increasePerBucket := float64(0.05)
	linkDifficulty := uint(2)
	rootAdditionalBuckets := uint(120)
	headAdditionalBuckets := uint(20)

	basePath := "./test_files"
	genesisFileNameWithExt := "genesis.json"
	defer func() {
		os.RemoveAll(basePath)
	}()

	chainFileName := "root"
	chainFileExt := "json"

	app := NewChainApplication(basePath, genesisFileNameWithExt, chainFileName, chainFileExt)
	err := app.Init(miningValue, baseDifficulty, increasePerBucket, linkDifficulty, rootAdditionalBuckets, headAdditionalBuckets)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRootChain, err := app.Retrieve()
	if err != nil {
		retRootChain.Root().Mining()
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRootAtIndex, err := app.RetrieveAtIndex(0)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(retRootChain, retRootAtIndex) {
		t.Errorf("the chains were expected to be the same")
		return
	}
}
