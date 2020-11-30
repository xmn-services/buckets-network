package bundles

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

/*
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
*/
func TestRestAPI_Success(t *testing.T) {
	basePath := "./test_files"

	// identity:
	name := "roger"
	password := "this-is-roger-pass"
	seed := "this is the seed of roger"
	rootDir := filepath.Join(basePath, "roger")

	// blockchain init:
	miningValue := uint8(0)
	baseDifficulty := uint(1)
	increasePerBucket := float64(0.05)
	linkDifficulty := uint(2)
	rootAdditionalBuckets := uint(120)
	headAdditionalBuckets := uint(20)

	// rest server:
	port := uint(7854)
	waitPeriod := time.Duration(15 * time.Second)
	maxUploadFileSize := int64(1024 * 1024 * 10)
	cmdApp := CreateCommandApplicationForTests()
	serverApp := NewRestAPIServer(cmdApp, maxUploadFileSize, waitPeriod, port)

	// create the base bucket folder:
	bucketDirPath := filepath.Join(basePath, "source_buckets")
	os.MkdirAll(bucketDirPath, os.ModePerm)

	// first bucket:
	firstBucketPath := filepath.Join(bucketDirPath, "first")
	os.MkdirAll(firstBucketPath, os.ModePerm)
	ioutil.WriteFile(filepath.Join(firstBucketPath, "first.data"), []byte("first bucket first data"), os.ModePerm)
	ioutil.WriteFile(filepath.Join(firstBucketPath, "second.data"), []byte("first bucket second data"), os.ModePerm)
	ioutil.WriteFile(filepath.Join(firstBucketPath, "third.data"), []byte("first bucket third data"), os.ModePerm)

	// second bucket:
	secondBucketPath := filepath.Join(bucketDirPath, "second")
	os.MkdirAll(secondBucketPath, os.ModePerm)
	ioutil.WriteFile(filepath.Join(secondBucketPath, "first.data"), []byte("second bucket first data"), os.ModePerm)
	ioutil.WriteFile(filepath.Join(secondBucketPath, "second.data"), []byte("second bucket second data"), os.ModePerm)

	// third bucket:
	thirdBucketPath := filepath.Join(bucketDirPath, "third")
	os.MkdirAll(thirdBucketPath, os.ModePerm)
	ioutil.WriteFile(filepath.Join(thirdBucketPath, "first.data"), []byte("third bucket first data"), os.ModePerm)

	// start the server in a new go routine:
	go serverApp.Start()

	defer func() {
		// delete the files:
		os.RemoveAll(basePath)

		// stop the server after running the tests
		serverApp.Stop()
	}()

	// wait a bit to make sure the server is started before we execute the tests:
	time.Sleep(10 * time.Second)

	// local peer:
	localPeer, err := peer.NewBuilder().WithHost("127.0.0.1").WithPort(port).IsClear().Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create a new identity:
	client := NewRestAPIClient(localPeer)
	err = client.Current().NewIdentity(name, password, seed, rootDir)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// authenticate:
	identityApp, err := client.Current().Authenticate(name, seed, password)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// init the blockchain:
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

	// retrieve the blockchain:
	chain, err := client.Sub().Chain().Retrieve()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the chain at index:
	chainAtIndexZero, err := client.Sub().Chain().RetrieveAtIndex(0)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !chain.Hash().Compare(chainAtIndexZero.Hash()) {
		t.Errorf("the two (2) chains were expected to contain the same hash")
		return
	}

	// add first bucket:
	err = identityApp.Sub().Bucket().Add(firstBucketPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// add second bucket:
	err = identityApp.Sub().Bucket().Add(secondBucketPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// add third bucket:
	err = identityApp.Sub().Bucket().Add(thirdBucketPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the buckets:
	retBuckets, err := identityApp.Sub().Bucket().RetrieveAll()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retBuckets) != 3 {
		t.Errorf("%d buckets wered expected, %d returned", 3, len(retBuckets))
		return
	}

	// delete 1 bucket:
	deletedBucket := retBuckets[1]
	err = identityApp.Sub().Bucket().Delete(deletedBucket.Hash().String())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// mine the block:
	additionalBlockBuckets := uint(rand.Intn(5) + 1)
	err = identityApp.Sub().Chain().Block(additionalBlockBuckets)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// mine the link:
	additionalLinkBuckets := uint(rand.Intn(5) + 1)
	err = identityApp.Sub().Chain().Link(additionalLinkBuckets)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// make sure the 2 remaining buckets are in the last mined block:

	// retrieve the buckets and make sure they are the same as the ones that were in the block:

	// delete 1 bucket:

	// retrieve the buckets, make sure there is only 1 bucket remaining:
}
