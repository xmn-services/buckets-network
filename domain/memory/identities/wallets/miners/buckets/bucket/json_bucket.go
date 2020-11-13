package bucket

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
)

// JSONBucket represents a JSON bucket instance
type JSONBucket struct {
	Hash         string              `json:"hash"`
	Bucket       *buckets.JSONBucket `json:"bucket"`
	AbsolutePath string              `json:"absolute_path"`
	PrivateKey   string              `json:"pk"`
	CreatedOn    time.Time           `json:"created_on"`
}

func createJSONBucketFromBucket(bucket Bucket) *JSONBucket {
	bucketAdapter := buckets.NewAdapter()
	jsBucket := bucketAdapter.ToJSON(bucket.Bucket())

	pkAdapter := encryption.NewAdapter()
	pk := pkAdapter.ToEncoded(bucket.PrivateKey())

	hsh := bucket.Hash().String()
	absolutePath := bucket.AbsolutePath()
	createdOn := bucket.CreatedOn()
	return createJSONBucket(hsh, jsBucket, absolutePath, pk, createdOn)
}

func createJSONBucket(
	hash string,
	bucket *buckets.JSONBucket,
	absolutePath string,
	pk string,
	createdOn time.Time,
) *JSONBucket {
	out := JSONBucket{
		Hash:         hash,
		Bucket:       bucket,
		AbsolutePath: absolutePath,
		PrivateKey:   pk,
		CreatedOn:    createdOn,
	}

	return &out
}
