package buckets

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type bucket struct {
	immutable entities.Immutable
	files     []files.File
	mp        map[string]files.File
}

func createBucketFromJSON(ins *JSONBucket) (Bucket, error) {
	fileLst := []files.File{}
	filesAdapter := files.NewAdapter()
	for _, oneJSFile := range ins.Files {
		fil, err := filesAdapter.ToFile(oneJSFile)
		if err != nil {
			return nil, err
		}

		fileLst = append(fileLst, fil)
	}

	return NewBuilder().
		Create().
		WithFiles(fileLst).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createBucket(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Bucket {
	return createBucketInternally(immutable, files, mp)
}

func createBucketInternally(
	immutable entities.Immutable,
	files []files.File,
	mp map[string]files.File,
) Bucket {
	out := bucket{
		immutable: immutable,
		files:     files,
		mp:        mp,
	}

	return &out
}

// Hash returns the hash
func (obj *bucket) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Files returns the files
func (obj *bucket) Files() []files.File {
	return obj.files
}

// FileByPath returns the file by path
func (obj *bucket) FileByPath(path string) (files.File, error) {
	if file, ok := obj.mp[path]; ok {
		return file, nil
	}

	str := fmt.Sprintf("the file path (%s) is invalid", path)
	return nil, errors.New(str)
}

// FileChunkByHash returns the file and chunk by chunk hash
func (obj *bucket) FileChunkByHash(hash hash.Hash) (files.File, chunks.Chunk, error) {
	var file files.File
	var chunk chunks.Chunk
	for _, oneFile := range obj.mp {
		chk, err := oneFile.ChunkByHash(hash)
		if err != nil {
			continue
		}

		file = oneFile
		chunk = chk
		break
	}

	if file != nil && chunk != nil {
		return file, chunk, nil
	}

	str := fmt.Sprintf("the chunk (hash: %s) does not exists in any file", hash.String())
	return nil, nil, errors.New(str)
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

	insBucket := pr.(*bucket)
	obj.immutable = insBucket.immutable
	obj.files = insBucket.files
	obj.mp = insBucket.mp
	return nil
}
