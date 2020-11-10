package chunks

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type chunk struct {
	immutable   entities.Immutable
	sizeInBytes uint
	data        hash.Hash
}

func createChunkFromJSON(ins *JSONChunk) (Chunk, error) {
	hashAdapter := hash.NewAdapter()
	data, err := hashAdapter.FromString(ins.Data)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithSizeInBytes(ins.SizeInBytes).
		WithData(*data).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createChunk(
	immutable entities.Immutable,
	sizeInBytes uint,
	data hash.Hash,
) Chunk {
	out := chunk{
		immutable:   immutable,
		sizeInBytes: sizeInBytes,
		data:        data,
	}

	return &out
}

// Hash returns the hash
func (obj *chunk) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// SizeInBytes returns the sizeInBytes
func (obj *chunk) SizeInBytes() uint {
	return obj.sizeInBytes
}

// Data returns the data hash
func (obj *chunk) Data() hash.Hash {
	return obj.data
}

// CreatedOn returns the creation time
func (obj *chunk) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *chunk) MarshalJSON() ([]byte, error) {
	ins := createJSONChunkFromChunk(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *chunk) UnmarshalJSON(data []byte) error {
	ins := new(JSONChunk)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createChunkFromJSON(ins)
	if err != nil {
		return err
	}

	insChunk := pr.(*chunk)
	obj.immutable = insChunk.immutable
	obj.sizeInBytes = insChunk.sizeInBytes
	obj.data = insChunk.data
	return nil
}
