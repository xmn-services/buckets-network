package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type file struct {
	immutable    entities.Immutable
	relativePath string
	chunks       []chunks.Chunk
	mp           map[string]chunks.Chunk
}

func createFileFromJSON(ins *JSONFile) (File, error) {
	chks := []chunks.Chunk{}
	chkAdapter := chunks.NewAdapter()
	for _, oneJSChunk := range ins.Chunks {
		chk, err := chkAdapter.ToChunk(oneJSChunk)
		if err != nil {
			return nil, err
		}

		chks = append(chks, chk)
	}

	return NewBuilder().
		Create().
		WithRelativePath(ins.RelativePath).
		WithChunks(chks).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createFile(
	immutable entities.Immutable,
	relativePath string,
	chunks []chunks.Chunk,
	mp map[string]chunks.Chunk,
) File {
	out := file{
		immutable:    immutable,
		relativePath: relativePath,
		chunks:       chunks,
		mp:           mp,
	}

	return &out
}

// Hash returns the hash
func (obj *file) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// RelativePath returns the relativePath
func (obj *file) RelativePath() string {
	return obj.relativePath
}

// Chunks returns the chunks
func (obj *file) Chunks() []chunks.Chunk {
	return obj.chunks
}

// ChunkByHash returns the chunk by hash
func (obj *file) ChunkByHash(hash hash.Hash) (chunks.Chunk, error) {
	keyname := hash.String()
	if chk, ok := obj.mp[keyname]; ok {
		return chk, nil
	}

	str := fmt.Sprintf("the chunk hash (%s) is invalid", keyname)
	return nil, errors.New(str)
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
	ins := new(JSONFile)
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
	obj.mp = insFile.mp
	return nil
}
