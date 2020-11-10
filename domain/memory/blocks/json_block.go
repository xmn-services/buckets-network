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

	bucketAdapter := buckets.NewAdapter()
	lst := block.Buckets()
	buckets := []*buckets.JSONBucket{}
	for _, oneBucket := range lst {
		jsBucket := bucketAdapter.ToJSON(oneBucket)
		buckets = append(buckets, jsBucket)
	}

	additional := block.Additional()
	createdOn := block.CreatedOn()
	return createJSONBlock(gen, buckets, additional, createdOn)
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
