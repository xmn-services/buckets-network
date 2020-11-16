package miners

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/blocks"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/permanents"
)

// JSONMiner represents a JSON miner instance
type JSONMiner struct {
	ToTransact  *buckets.JSONBuckets    `json:"to_transact"`
	Queue       *buckets.JSONBuckets    `json:"queue"`
	Broadcasted *permanents.JSONBuckets `json:"broadcasted"`
	ToLink      *blocks.JSONBlocks      `json:"to_link"`
}

func createJSONMinerFromMiner(miner Miner) *JSONMiner {
	bucketsAdapter := buckets.NewAdapter()
	toTransact := bucketsAdapter.ToJSON(miner.ToTransact())
	queue := bucketsAdapter.ToJSON(miner.Queue())

	permanentBucketsAdapter := permanents.NewAdapter()
	broadcasted := permanentBucketsAdapter.ToJSON(miner.Broadcasted())

	blockAdapter := blocks.NewAdapter()
	toLink := blockAdapter.ToJSON(miner.ToLink())
	return createJSONMiner(toTransact, queue, broadcasted, toLink)
}

func createJSONMiner(
	toTransact *buckets.JSONBuckets,
	queue *buckets.JSONBuckets,
	broadcasted *permanents.JSONBuckets,
	toLink *blocks.JSONBlocks,
) *JSONMiner {
	out := JSONMiner{
		ToTransact:  toTransact,
		Queue:       queue,
		Broadcasted: broadcasted,
		ToLink:      toLink,
	}

	return &out
}
