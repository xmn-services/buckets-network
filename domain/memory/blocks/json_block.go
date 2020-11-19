package blocks

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/genesis"
)

// JSONBlock represents a JSON block instance
type JSONBlock struct {
	Genesis    *genesis.JSONGenesis  `json:"genesis"`
	Buckets    []*buckets.JSONBucket `json:"buckets"`
	Additional uint                  `json:"additional"`
	CreatedOn  time.Time             `json:"created_on"`
}

func createJSONBlockFromBlock(block Block) *JSONBlock {
	genAdapter := genesis.NewAdapter()
	gen := genAdapter.ToJSON(block.Genesis())

	jsonBuckets := []*buckets.JSONBucket{}
	if block.HasBuckets() {
		bucketAdapter := buckets.NewAdapter()
		lst := block.Buckets()
		for _, oneBucket := range lst {
			jsBucket := bucketAdapter.ToJSON(oneBucket)
			jsonBuckets = append(jsonBuckets, jsBucket)
		}
	}

	additional := block.Additional()
	createdOn := block.CreatedOn()
	return createJSONBlock(gen, jsonBuckets, additional, createdOn)
}

func createJSONBlock(
	gen *genesis.JSONGenesis,
	buckets []*buckets.JSONBucket,
	additional uint,
	createdOn time.Time,
) *JSONBlock {
	out := JSONBlock{
		Genesis:    gen,
		Buckets:    buckets,
		Additional: additional,
		CreatedOn:  createdOn,
	}

	return &out
}
