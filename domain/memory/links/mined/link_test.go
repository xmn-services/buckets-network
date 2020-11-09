package mined

import (
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	"github.com/xmn-services/buckets-network/domain/memory/transactions"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	bucketHash, _ := hashAdapter.FromBytes([]byte("lets say this is the bucket hash"))

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(0.00001)
	linkDiff := uint(1)
	gen := genesis.CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff)

	// transaction:
	executesOnTrigger := true
	amountPubKeyInRing := uint(20)
	trx := transactions.CreateTransactionForTests(amountPubKeyInRing, executesOnTrigger, *bucketHash)

	// block:
	additional := uint(0)
	trxList := []transactions.Transaction{
		trx,
	}

	nextBlock := blocks.CreateBlockForTests(gen, additional, trxList)

	// link:
	index := uint(2)
	prevLink, _ := hashAdapter.FromBytes([]byte("prev link hash"))
	link := links.CreateLinkForTests(*prevLink, nextBlock, index)

	// mined link:
	mining := "232"
	minedLink := CreateLinkForTests(link, mining)

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
