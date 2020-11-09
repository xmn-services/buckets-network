package links

import (
	"encoding/json"
	"testing"

	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
	"github.com/xmn-services/buckets-network/domain/memory/transactions"
	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestLink_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	bucketHash, _ := hashAdapter.FromBytes([]byte("lets say this is the bucket hash"))

	// transaction:
	executesOnTrigger := true
	amountPubKeyInRing := uint(20)
	trxIns := transactions.CreateTransactionForTests(amountPubKeyInRing, executesOnTrigger, *bucketHash)

	// transactions:
	trx := []transactions.Transaction{
		trxIns,
	}

	// genesis:
	blockDiffBase := uint(1)
	blockDiffIncreasePerTrx := float64(1.0)
	linkDiff := uint(1)
	genesisIns := genesis.CreateGenesisForTests(blockDiffBase, blockDiffIncreasePerTrx, linkDiff)

	// block:
	additional := uint(0)
	blockIns := blocks.CreateBlockForTests(genesisIns, additional, trx)

	// link:
	index := uint(5)
	prevLink, _ := hashAdapter.FromBytes([]byte("this is the prev link"))
	linkIns := CreateLinkForTests(*prevLink, blockIns, index)

	// repository and service:
	genService, repository, service := CreateRepositoryServiceForTests()
	err := service.Save(linkIns)
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

	retTrx, err := repository.Retrieve(linkIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !linkIns.Hash().Compare(retTrx.Hash()) {
		t.Errorf("the hash is invalid")
		return
	}

	js, err := json.Marshal(linkIns)
	if err != nil {
		t.Errorf("the Link instance could not be converted to JSON: %s", err.Error())
		return
	}

	ins := new(link)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the JSON instance could not be converted to a Link instance: %s", err.Error())
		return
	}

	// compare:
	TestCompare(t, ins, linkIns)
}
