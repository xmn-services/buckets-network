package files

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
)

// JSONFile represents a JSON file instance
type JSONFile struct {
	RelativePath string              `json:"relative_path"`
	Chunks       []*chunks.JSONChunk `json:"chunks"`
	CreatedOn    time.Time           `json:"created_on"`
}

func createJSONFileFromFile(file File) *JSONFile {
	relativePath := file.RelativePath()

	chunkAdapter := chunks.NewAdapter()
	jsonChunks := []*chunks.JSONChunk{}
	chunks := file.Chunks()
	for _, oneChunk := range chunks {
		chk := chunkAdapter.ToJSON(oneChunk)
		jsonChunks = append(jsonChunks, chk)
	}

	createdOn := file.CreatedOn()
	return createJSONFile(relativePath, jsonChunks, createdOn)
}

func createJSONFile(
	relativePath string,
	chunks []*chunks.JSONChunk,
	createdOn time.Time,
) *JSONFile {
	out := JSONFile{
		RelativePath: relativePath,
		Chunks:       chunks,
		CreatedOn:    createdOn,
	}

	return &out
}
