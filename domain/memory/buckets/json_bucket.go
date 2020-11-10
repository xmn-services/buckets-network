package buckets

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
)

// JSONBucket represents a JSON bucket instance
type JSONBucket struct {
	Files     []*files.JSONFile `json:"files"`
	CreatedOn time.Time         `json:"created_on"`
}

func createJSONBucketFromBucket(bucket Bucket) *JSONBucket {
	chunkAdapter := files.NewAdapter()
	jsonFiles := []*files.JSONFile{}
	files := bucket.Files()
	for _, oneChunk := range files {
		chk := chunkAdapter.ToJSON(oneChunk)
		jsonFiles = append(jsonFiles, chk)
	}

	createdOn := bucket.CreatedOn()
	return createJSONBucket(jsonFiles, createdOn)
}

func createJSONBucket(
	files []*files.JSONFile,
	createdOn time.Time,
) *JSONBucket {
	out := JSONBucket{
		Files:     files,
		CreatedOn: createdOn,
	}

	return &out
}
