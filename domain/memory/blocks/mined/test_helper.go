package mined

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	transfer_block_mined "github.com/xmn-services/buckets-network/domain/transfers/blocks/mined"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CreateBlockForTests creates a mined block instance for tests
func CreateBlockForTests(blk blocks.Block, mining string) Block {
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithBlock(blk).WithMining(mining).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// CreateBlockForTestsWithoutParams create block for tests without params
func CreateBlockForTestsWithoutParams() (genesis.Genesis, Block) {
	hashAdapter := hash.NewAdapter()
	firstHash, _ := hashAdapter.FromBytes([]byte("this is the first hash"))
	secondHash, _ := hashAdapter.FromBytes([]byte("this is the second hash"))

	firstChunks := []chunks.Chunk{
		chunks.CreateChunkForTests(uint(345234), *firstHash),
	}

	secondChunks := []chunks.Chunk{
		chunks.CreateChunkForTests(uint(2345234), *secondHash),
	}

	firstRelativePath := "/first/is/relative/path"
	secondRelativePath := "/second/is/relative/path"

	files := []files.File{
		files.CreateFileForTests(firstRelativePath, firstChunks),
		files.CreateFileForTests(secondRelativePath, secondChunks),
	}

	bucketIns := buckets.CreateBucketForTests(files)

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(1.0)
	linkDiff := uint(1)
	genesisIns := genesis.CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff)

	// block:
	additional := uint(0)
	blockIns := blocks.CreateBlockForTests(genesisIns, additional, []buckets.Bucket{
		bucketIns,
	})

	// mined block:
	mining := "sdfgfgsdf"
	return genesisIns, CreateBlockForTests(blockIns, mining)
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (genesis.Service, Repository, Service) {
	genesisService, blockRepository, blockService := blocks.CreateRepositoryServiceForTests()
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	trRepository := transfer_block_mined.NewRepository(fileRepositoryService)
	trService := transfer_block_mined.NewService(fileRepositoryService)
	repository := NewRepository(blockRepository, trRepository)
	service := NewService(repository, blockService, trService)
	return genesisService, repository, service
}

// TestCompare compare two block instances
func TestCompare(t *testing.T, first Block, second Block) {
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
