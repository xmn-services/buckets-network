package buckets

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type buckets struct {
	lst []bucket.Bucket
	mp  map[string]bucket.Bucket
}

func createBucketsFromJSON(ins *JSONBuckets) (Buckets, error) {
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
		WithBuckets(buckets).
		Now()
}

func crateBuckets(
	lst []bucket.Bucket,
	mp map[string]bucket.Bucket,
) Buckets {
	out := buckets{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All return all buckets
func (obj *buckets) All() []bucket.Bucket {
	return obj.lst
}

// Fetch fetches a bucket by hash
func (obj *buckets) Fetch(hash hash.Hash) (bucket.Bucket, error) {
	keyname := hash.String()
	if bucket, ok := obj.mp[keyname]; ok {
		return bucket, nil
	}

	str := fmt.Sprintf("the bucket (hash: %s) does not exists", keyname)
	return nil, errors.New(str)
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
	obj.lst = insBucket.lst
	obj.mp = insBucket.mp
	return nil
}
