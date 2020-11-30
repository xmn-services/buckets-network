package contents

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/libs/hash"
)

type content struct {
	bucket hash.Hash
	file   hash.Hash
	chunk  hash.Hash
}

func createContentFromJSON(ins *JSONContent) (Content, error) {
	hashAdapter := hash.NewAdapter()
	bucket, err := hashAdapter.FromString(ins.Bucket)
	if err != nil {
		return nil, err
	}

	file, err := hashAdapter.FromString(ins.File)
	if err != nil {
		return nil, err
	}

	chunk, err := hashAdapter.FromString(ins.Chunk)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithBucket(*bucket).
		WithFile(*file).
		WithChunk(*chunk).
		Now()
}

func createContent(
	bucket hash.Hash,
	file hash.Hash,
	chunk hash.Hash,
) Content {
	out := content{
		bucket: bucket,
		file:   file,
		chunk:  chunk,
	}

	return &out
}

// Bucket returns the bucket hash
func (obj *content) Bucket() hash.Hash {
	return obj.bucket
}

// File returns the file hash
func (obj *content) File() hash.Hash {
	return obj.file
}

// Chunk returns the chunk hash
func (obj *content) Chunk() hash.Hash {
	return obj.chunk
}

// MarshalJSON converts the instance to JSON
func (obj *content) MarshalJSON() ([]byte, error) {
	ins := createJSONContentFromContent(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *content) UnmarshalJSON(data []byte) error {
	ins := new(JSONContent)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createContentFromJSON(ins)
	if err != nil {
		return err
	}

	insContent := pr.(*content)
	obj.bucket = insContent.bucket
	obj.file = insContent.file
	obj.chunk = insContent.chunk
	return nil
}
