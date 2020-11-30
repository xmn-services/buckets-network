package contents

import (
	"path/filepath"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	fileRepository file.Repository
}

func createRepository(
	fileRepository file.Repository,
) Repository {
	out := repository{
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a chunk's data from bucket
func (app *repository) Retrieve(bucketHash hash.Hash, fileHash hash.Hash, chunkHash hash.Hash) ([]byte, error) {
	path := filepath.Join(bucketHash.String(), fileHash.String(), chunkHash.String())
	return app.fileRepository.Retrieve(path)
}
