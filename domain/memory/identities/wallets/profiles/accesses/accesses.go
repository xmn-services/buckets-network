package accesses

import (
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses/access"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type accesses struct {
	lst []access.Access
	mp  map[string]access.Access
}

func createAccesses(
	lst []access.Access,
	mp map[string]access.Access,
) Accesses {
	out := accesses{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All returns the accesses
func (obj *accesses) All() []access.Access {
	return obj.lst
}

// Add adds an access
func (obj *accesses) Add(access access.Access) error {
	keyname := access.Bucket().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the bucket access (hash: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, access)
	obj.mp[keyname] = access
	return nil
}

// Fetch fetches an access by bucket hash
func (obj *accesses) Fetch(bucket hash.Hash) (access.Access, error) {
	keyname := bucket.String()
	if access, ok := obj.mp[keyname]; ok {
		return access, nil
	}

	str := fmt.Sprintf("the bucket access (hash: %s) does NOT exists", keyname)
	return nil, errors.New(str)
}

// Delete deletes an access
func (obj *accesses) Delete(bucket hash.Hash) error {
	keyname := bucket.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the bucket access (hash: %s) does not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneAccess := range obj.lst {
		if oneAccess.Bucket().Compare(bucket) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}
