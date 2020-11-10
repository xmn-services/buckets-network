package chunks

import (
	transfer_chunk "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks"
)

type adapter struct {
	trBuilder transfer_chunk.Builder
}

func createAdapter(
	trBuilder transfer_chunk.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a chunk to a transfer chunk instance
func (app *adapter) ToTransfer(chunk Chunk) (transfer_chunk.Chunk, error) {
	hash := chunk.Hash()
	sizeInBytes := chunk.SizeInBytes()
	data := chunk.Data()
	createdOn := chunk.CreatedOn()
	return app.trBuilder.Create().WithHash(hash).WithSizeInBytes(sizeInBytes).WithData(data).CreatedOn(createdOn).Now()
}

// ToJSON converts a chunk to JSON
func (app *adapter) ToJSON(chunk Chunk) *JSONChunk {
	return createJSONChunkFromChunk(chunk)
}

// ToChunk converts JSON to chunk
func (app *adapter) ToChunk(ins *JSONChunk) (Chunk, error) {
	return createChunkFromJSON(ins)
}
