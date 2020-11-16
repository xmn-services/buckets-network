package miners

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
)

// JSONMiner represents a JSON miner instance
type JSONMiner struct {
	Hash          string                  `json:"hash"`
	ToTransact    *buckets.JSONBuckets    `json:"to_transact"`
	Queue         *buckets.JSONBuckets    `json:"queue"`
	Broadcasted   *permanents.JSONBuckets `json:"broadcasted"`
	ToLink        *blocks.JSONBlocks      `json:"to_link"`
	CreatedOn     time.Time               `json:"created_on"`
	LastUpdatedOn time.Time               `json:"last_updated_on"`
}

func createJSONMinerFromMiner(miner Miner) *JSONMiner {
	bucketsAdapter := buckets.NewAdapter()
	toTransact := bucketsAdapter.ToJSON(miner.ToTransact())
	queue := bucketsAdapter.ToJSON(miner.Queue())

	permanentBucketsAdapter := permanents.NewAdapter()
	broadcasted := permanentBucketsAdapter.ToJSON(miner.Broadcasted())

	blockAdapter := blocks.NewAdapter()
	toLink := blockAdapter.ToJSON(miner.ToLink())

	hsh := miner.Hash().String()
	createdOn := miner.CreatedOn()
	lastUpdatedOn := miner.LastUpdatedOn()
	return createJSONMiner(hsh, toTransact, queue, broadcasted, toLink, createdOn, lastUpdatedOn)
}

func createJSONMiner(
	hash string,
	toTransact *buckets.JSONBuckets,
	queue *buckets.JSONBuckets,
	broadcasted *permanents.JSONBuckets,
	toLink *blocks.JSONBlocks,
	createdOn time.Time,
	lastUpdatedOn time.Time,
) *JSONMiner {
	out := JSONMiner{
		Hash:          hash,
		ToTransact:    toTransact,
		Queue:         queue,
		Broadcasted:   broadcasted,
		ToLink:        toLink,
		CreatedOn:     createdOn,
		LastUpdatedOn: lastUpdatedOn,
	}

	return &out
}
