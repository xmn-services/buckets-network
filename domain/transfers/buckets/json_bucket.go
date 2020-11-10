package buckets

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type jsonBucket struct {
	Hash      string                `json:"hash"`
	Files     *hashtree.JSONCompact `json:"files"`
	Amount    uint                  `json:"amount"`
	CreatedOn time.Time             `json:"created_on"`
}

func createJSONBucketFromBucket(ins Bucket) *jsonBucket {
	hash := ins.Hash().String()
	files := hashtree.NewAdapter().ToJSON(ins.Files().Compact())
	amount := ins.Amount()
	createdOn := ins.CreatedOn()
	return createJSONBucket(hash, files, amount, createdOn)
}

func createJSONBucket(
	hash string,
	files *hashtree.JSONCompact,
	amount uint,
	createdOn time.Time,
) *jsonBucket {
	out := jsonBucket{
		Hash:      hash,
		Files:     files,
		Amount:    amount,
		CreatedOn: createdOn,
	}

	return &out
}
