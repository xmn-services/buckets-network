package buckets

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type bucket struct {
	immutable entities.Immutable
	files     hashtree.HashTree
	amount    uint
}

func createBucketFromJSON(ins *jsonBucket) (Bucket, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	compact, err := hashtree.NewAdapter().FromJSON(ins.Files)
	if err != nil {
		return nil, err
	}

	files, err := compact.Leaves().HashTree()
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithFiles(files).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBucket(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Bucket {
	return createBucketInternally(immutable, files, amount)
}

func createBucketInternally(
	immutable entities.Immutable,
	files hashtree.HashTree,
	amount uint,
) Bucket {
	out := bucket{
		immutable: immutable,
		files:     files,
		amount:    amount,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files return the files
func (obj *bucket) Files() hashtree.HashTree {
	return obj.files
}

// Amount returns the amount
func (obj *bucket) Amount() uint {
	return obj.amount
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
	ins := new(jsonBucket)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBucketFromJSON(ins)
	if err != nil {
		return err
	}

	insBucket := pr.(*bucket)
	obj.immutable = insBucket.immutable
	obj.files = insBucket.files
	obj.amount = insBucket.amount
	return nil
}
