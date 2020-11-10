package chunks

import (
	"time"
)

// JSONChunk represents a JSON chunk instance
type JSONChunk struct {
	SizeInBytes uint      `json:"size_in_bytes"`
	Data        string    `json:"data"`
	CreatedOn   time.Time `json:"created_on"`
}

func createJSONChunkFromChunk(chunk Chunk) *JSONChunk {
	sizeInBytes := chunk.SizeInBytes()
	data := chunk.Data().String()
	createdOn := chunk.CreatedOn()
	return createJSONChunk(sizeInBytes, data, createdOn)
}

func createJSONChunk(
	sizeInBytes uint,
	data string,
	createdOn time.Time,
) *JSONChunk {
	out := JSONChunk{
		SizeInBytes: sizeInBytes,
		Data:        data,
		CreatedOn:   createdOn,
	}

	return &out
}
