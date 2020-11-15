package permanents

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
)

// JSONBuckets represents a JSON buckets instance
type JSONBuckets struct {
	Hash          string               `json:"hash"`
	Buckets       []*bucket.JSONBucket `json:"bucket"`
	CreatedOn     time.Time            `json:"created_on"`
	LastUpdatedOn time.Time            `json:"last_updated_on"`
}

func createJSONBucketsFromBuckets(ins Buckets) *JSONBuckets {
	bucketAdapter := bucket.NewAdapter()

	jsonBuckets := []*bucket.JSONBucket{}
	bckets := ins.All()
	for _, oneBucket := range bckets {
		jsonBucket := bucketAdapter.ToJSON(oneBucket)
		jsonBuckets = append(jsonBuckets, jsonBucket)
	}

	hsh := ins.Hash().String()
	createdOn := ins.CreatedOn()
	lastUpdatedOn := ins.LastUpdatedOn()
	return createJSONBuckets(hsh, jsonBuckets, createdOn, lastUpdatedOn)
}

func createJSONBuckets(
	hash string,
	buckets []*bucket.JSONBucket,
	createdOn time.Time,
	lastUpdatedOn time.Time,
) *JSONBuckets {
	out := JSONBuckets{
		Hash:          hash,
		Buckets:       buckets,
		CreatedOn:     createdOn,
		LastUpdatedOn: lastUpdatedOn,
	}

	return &out
}
