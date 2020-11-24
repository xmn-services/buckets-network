package bundles

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestInit_Success(t *testing.T) {
	miningValue := uint8(0)
	baseDifficulty := uint(1)
	increasePerBucket := float64(0.05)
	linkDifficulty := uint(2)
	rootAdditionalBuckets := uint(120)
	headAdditionalBuckets := uint(20)

	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// identity:
	name := "roger"
	password := "this-is-roger-pass"
	seed := "this is the seed of roger"
	rootDir := filepath.Join(basePath, "roger")

	app := CreateCommandApplicationForTests()
	err := app.Current().NewIdentity(name, password, seed, rootDir)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	identityApp, err := app.Current().Authenticate(name, password, seed)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = identityApp.Sub().Chain().Init(
		miningValue,
		baseDifficulty,
		increasePerBucket,
		linkDifficulty,
		rootAdditionalBuckets,
		headAdditionalBuckets,
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRootChain, err := app.Sub().Chain().Retrieve()
	if err != nil {
		retRootChain.Root().Mining()
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retRootAtIndex, err := app.Sub().Chain().RetrieveAtIndex(0)
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

func TestRestAPI_Success(t *testing.T) {
	port := uint(80)
	waitPeriod := time.Duration(15 * time.Second)
	maxUploadFileSize := int64(1024 * 1024 * 10)
	cmdApp := CreateCommandApplicationForTests()
	serverApp := NewRestAPIServer(cmdApp, maxUploadFileSize, waitPeriod, port)
	err := serverApp.Start()
	defer func() {
		serverApp.Stop()
	}()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
