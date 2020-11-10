package blocks

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type jsonBlock struct {
	Hash       string                `json:"hash"`
	Buckets    *hashtree.JSONCompact `json:"buckets"`
	Amount     uint                  `json:"amount"`
	Additional uint                  `json:"additional"`
	CreatedOn  time.Time             `json:"created_on"`
}

func createJSONBlockFromBlock(ins Block) *jsonBlock {
	hash := ins.Hash().String()
	buckets := hashtree.NewAdapter().ToJSON(ins.Buckets().Compact())
	amount := ins.Amount()
	additional := ins.Additional()
	createdOn := ins.CreatedOn()
	return createJSONBlock(hash, buckets, amount, additional, createdOn)
}

func createJSONBlock(
	hash string,
	buckets *hashtree.JSONCompact,
	amount uint,
	additional uint,
	createdOn time.Time,
) *jsonBlock {
	out := jsonBlock{
		Hash:       hash,
		Buckets:    buckets,
		Amount:     amount,
		Additional: additional,
		CreatedOn:  createdOn,
	}

	return &out
}
