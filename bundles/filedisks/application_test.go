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

func TestPeer_isClear_Success(t *testing.T) {
	firstHost := "78.98.32.34"
	firstPort := uint(8080)

	secondHost := "sdsdgfg"
	secondPort := uint(443)

	basePath := "./test_files"
	fileNameWithExt := "peers.json"
	defer func() {
		os.RemoveAll(basePath)
	}()

	peerApp := NewPeerApplication(basePath, fileNameWithExt)
	err := peerApp.SaveClear(firstHost, firstPort)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstRetPeers, err := peerApp.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstAmount := len(firstRetPeers.All())
	if firstAmount != 1 {
		t.Errorf("%d peers were expected, %d returned", firstAmount, len(firstRetPeers.All()))
		return
	}

	err = peerApp.SaveOnion(secondHost, secondPort)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secondRetPeers, err := peerApp.Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secondAmount := len(secondRetPeers.All())
	if secondAmount != 2 {
		t.Errorf("%d peers were expected, %d returned", secondAmount, len(secondRetPeers.All()))
		return
	}
}
