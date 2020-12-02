package accesses

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type accesses struct {
	lst []hash.Hash
	mp  map[string]hash.Hash
}

func createAccessesFromJSON(ins *JSONAccesses) (Accesses, error) {
	hashAdapter := hash.NewAdapter()
	lst := []hash.Hash{}
	for _, oneHash := range ins.List {
		hsh, err := hashAdapter.FromString(oneHash)
		if err != nil {
			return nil, err
		}

		lst = append(lst, *hsh)
	}

	return NewBuilder().
		Create().
		WithList(lst).
		Now()
}

func createAccesses(
	lst []hash.Hash,
	mp map[string]hash.Hash,
) Accesses {
	out := accesses{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All returns the bucket hashes
func (obj *accesses) All() []hash.Hash {
	return obj.lst
}

// Add adds a bucket
func (obj *accesses) Add(bucket buckets.Bucket) error {
	keyname := bucket.Hash().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the bucket (hash: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, bucket.Hash())
	obj.mp[keyname] = bucket.Hash()
	return nil
}

// Fetch fetches a bucket by hash
func (obj *accesses) Fetch(bucket hash.Hash) (*hash.Hash, error) {
	keyname := bucket.String()
	if hsh, ok := obj.mp[keyname]; ok {
		return &hsh, nil
	}

	str := fmt.Sprintf("the bucket (hash: %s) does NOT exists", keyname)
	return nil, errors.New(str)
}

// Delete deletes a bucket
func (obj *accesses) Delete(bucket hash.Hash) error {
	keyname := bucket.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the bucket (hash: %s) does not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneHash := range obj.lst {
		if oneHash.Compare(bucket) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *accesses) MarshalJSON() ([]byte, error) {
	ins := createJSONAccessesFromAccesses(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *accesses) UnmarshalJSON(data []byte) error {
	ins := new(JSONAccesses)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createAccessesFromJSON(ins)
	if err != nil {
		return err
	}

	insAccesses := pr.(*accesses)
	obj.lst = insAccesses.lst
	obj.mp = insAccesses.mp
	return nil
}
