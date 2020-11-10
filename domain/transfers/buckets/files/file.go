package files

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type file struct {
	immutable    entities.Immutable
	relativePath string
	chunks       hashtree.HashTree
	amount       uint
}

func createFileFromJSON(ins *jsonFile) (File, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	compact, err := hashtree.NewAdapter().FromJSON(ins.Chunks)
	if err != nil {
		return nil, err
	}

	chunks, err := compact.Leaves().HashTree()
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithChunks(chunks).
		WithRelativePath(ins.RelativePath).
		WithAmount(ins.Amount).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createFile(
	immutable entities.Immutable,
	relativePath string,
	chunks hashtree.HashTree,
	amount uint,
) File {
	out := file{
		immutable:    immutable,
		relativePath: relativePath,
		chunks:       chunks,
		amount:       amount,
	}

	return &out
}

// Hash returns the hash
func (obj *file) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// RelativePath returns the relative path
func (obj *file) RelativePath() string {
	return obj.relativePath
}

// Chunks returns the chunks
func (obj *file) Chunks() hashtree.HashTree {
	return obj.chunks
}

// Amount returns the files amount
func (obj *file) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *file) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *file) MarshalJSON() ([]byte, error) {
	ins := createJSONFileFromFile(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *file) UnmarshalJSON(data []byte) error {
	ins := new(jsonFile)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createFileFromJSON(ins)
	if err != nil {
		return err
	}

	insFile := pr.(*file)
	obj.immutable = insFile.immutable
	obj.relativePath = insFile.relativePath
	obj.chunks = insFile.chunks
	obj.amount = insFile.amount
	return nil
}
