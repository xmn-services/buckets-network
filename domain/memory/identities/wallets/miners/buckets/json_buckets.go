package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
)

// JSONBuckets represents a JSON buckets instance
type JSONBuckets struct {
	Buckets []*bucket.JSONBucket `json:"buckets"`
}

func createJSONBucketsFromBuckets(ins Buckets) *JSONBuckets {
	bucketAdapter := bucket.NewAdapter()

	jsonBuckets := []*bucket.JSONBucket{}
	bckets := ins.All()
	for _, oneBucket := range bckets {
		jsonBucket := bucketAdapter.ToJSON(oneBucket)
		jsonBuckets = append(jsonBuckets, jsonBucket)
	}

	return createJSONBuckets(jsonBuckets)
}

func createJSONBuckets(
	buckets []*bucket.JSONBucket,
) *JSONBuckets {
	out := JSONBuckets{
		Buckets: buckets,
	}

	return &out
}
