package bucket

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type bucket struct {
	immutable    entities.Immutable
	bucket       buckets.Bucket
	absolutePath string
	pk           encryption.PrivateKey
}

func createBucketFromJSON(ins *JSONBucket) (Bucket, error) {
	bucketAdapter := buckets.NewAdapter()
	bucket, err := bucketAdapter.ToBucket(ins.Bucket)
	if err != nil {
		return nil, err
	}

	pkAdapter := encryption.NewAdapter()
	pk, err := pkAdapter.FromEncoded(ins.PrivateKey)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithAbsolutePath(ins.AbsolutePath).
		WithPrivateKey(pk).
		WithBucket(bucket).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBucket(
	immutable entities.Immutable,
	bket buckets.Bucket,
	absolutePath string,
	pk encryption.PrivateKey,
) Bucket {
	out := bucket{
		immutable:    immutable,
		bucket:       bket,
		absolutePath: absolutePath,
		pk:           pk,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Bucket returns the Bucket
func (obj *bucket) Bucket() buckets.Bucket {
	return obj.bucket
}

// AbsolutePath returns the absolutePath
func (obj *bucket) AbsolutePath() string {
	return obj.absolutePath
}

// PrivateKey returns the privateKey
func (obj *bucket) PrivateKey() encryption.PrivateKey {
	return obj.pk
}

// CreatedOn returns the creation time
func (obj *bucket) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *bucket) MarshalJSON() ([]byte, error) {
	ins := createJSONBucketFromBucket(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *bucket) UnmarshalJSON(data []byte) error {
	ins := new(JSONBucket)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBucketFromJSON(ins)
	if err != nil {
		return err
	}

	insBlock := pr.(*bucket)
	obj.immutable = insBlock.immutable
	obj.bucket = insBlock.bucket
	obj.absolutePath = insBlock.absolutePath
	obj.pk = insBlock.pk
	return nil
}
