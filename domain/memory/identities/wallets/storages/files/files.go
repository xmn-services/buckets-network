package files

import (
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/libs/hash"
)

type files struct {
	lst []hash.Hash
	mp  map[string]hash.Hash
}

func createFilesFromJSON(ins *JSONFiles) (Files, error) {
	files := []hash.Hash{}
	hashAdapter := hash.NewAdapter()
	for _, oneHashStr := range ins.Files {
		hsh, err := hashAdapter.FromString(oneHashStr)
		if err != nil {
			return nil, err
		}

		files = append(files, *hsh)
	}

	return NewBuilder().
		Create().
		WithFiles(files).
		Now()
}

func createFiles(
	lst []hash.Hash,
	mp map[string]hash.Hash,
) Files {
	out := files{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All return all file hashes
func (obj *files) All() []hash.Hash {
	return obj.lst
}

// Exists returns true if the file hash already exists, false otherwise
func (obj *files) Exists(hash hash.Hash) bool {
	keyname := hash.String()
	if _, ok := obj.mp[keyname]; ok {
		return true
	}

	return false
}

// Add adds a file hash to the list
func (obj *files) Add(hash hash.Hash) error {
	keyname := hash.String()
	if obj.Exists(hash) {
		str := fmt.Sprintf("the file (hash: %s) already exists and therefore cannot be added", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, hash)
	obj.mp[keyname] = hash
	return nil
}

// Delete deletes a file hash to the list
func (obj *files) Delete(hash hash.Hash) error {
	keyname := hash.String()
	if !obj.Exists(hash) {
		str := fmt.Sprintf("the file (hash: %s) do not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneHash := range obj.lst {
		if oneHash.Compare(hash) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}
