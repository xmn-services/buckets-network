package mined

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
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
	blockIns := mined_blocks.CreateBlockForTests(genesisIns, additional, []buckets.Bucket{
		bucketIns,
	})

	// mined block:
	mining := "this is a mining"
	minedBlockIns := mined_mined_blocks.CreateBlockForTests(blockIns, mining)

	// link:
	index := uint(5)
	prevLink, _ := hashAdapter.FromBytes([]byte("this is the prev link"))
	link := links.CreateLinkForTests(*prevLink, minedBlockIns, index)

	// mined link:
	linkMining := "232"
	minedLink := CreateLinkForTests(link, linkMining)

	// encode then decode:
	adapter := NewAdapter()
	encoded, err := adapter.Encode(minedLink)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	decoded, err := adapter.Decode(encoded)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reEncoded, err := adapter.Encode(decoded)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if encoded != reEncoded {
		t.Errorf("the encoding and decoding process failed to work")
		return
	}
}
