package mined

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestBlock_Success(t *testing.T) {
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
	minedBlockIns := CreateBlockForTests(blockIns, mining)

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
