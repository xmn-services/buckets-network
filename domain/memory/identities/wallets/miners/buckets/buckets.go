package buckets

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type buckets struct {
	mutable entities.Mutable
	lst     []bucket.Bucket
	mp      map[string]bucket.Bucket
}

func createBucketsFromJSON(ins *JSONBuckets) (Buckets, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	buckets := []bucket.Bucket{}
	bucketAdapter := bucket.NewAdapter()
	for _, oneJS := range ins.Buckets {
		bucket, err := bucketAdapter.ToBucket(oneJS)
		if err != nil {
			return nil, err
		}

		buckets = append(buckets, bucket)
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithBuckets(buckets).
		CreatedOn(ins.CreatedOn).
		LastUpdatedOn(ins.LastUpdatedOn).
		Now()
}

func crateBuckets(
	mutable entities.Mutable,
	lst []bucket.Bucket,
	mp map[string]bucket.Bucket,
) Buckets {
	out := buckets{
		mutable: mutable,
		lst:     lst,
		mp:      mp,
	}

	return &out
}

// Hash returns the hash
func (obj *buckets) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All return all buckets
func (obj *buckets) All() []bucket.Bucket {
	return obj.lst
}

// Add adds a bucket to the list
func (obj *buckets) Add(bucket bucket.Bucket) error {
	keyname := bucket.Hash().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the bucket (hash: %s) cannot be added because it already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, bucket)
	obj.mp[keyname] = bucket
	return nil
}

// Delete deletes a bucket from the list
func (obj *buckets) Delete(hash hash.Hash) error {
	keyname := hash.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the bucket (hash: %s) cannot be deleted because it does NOT exists", keyname)
		return errors.New(str)
	}

	for index, oneBucket := range obj.lst {
		if oneBucket.Hash().Compare(hash) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// CreatedOn returns the creation time
func (obj *buckets) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *buckets) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *buckets) MarshalJSON() ([]byte, error) {
	ins := createJSONBucketsFromBuckets(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *buckets) UnmarshalJSON(data []byte) error {
	ins := new(JSONBuckets)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBucketsFromJSON(ins)
	if err != nil {
		return err
	}

	insBucket := pr.(*buckets)
	obj.mutable = insBucket.mutable
	obj.lst = insBucket.lst
	obj.mp = insBucket.mp
	return nil
}
