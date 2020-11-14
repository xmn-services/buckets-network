package data

import (
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type repository struct {
	fileRepository file.Repository
}

func createRepository(fileRepository file.Repository) Repository {
	out := repository{
		fileRepository: fileRepository,
	}

	return &out
}

// Exists returns true if the data exists, false otherwise
func (app *repository) Exists(hash hash.Hash) bool {
	return app.fileRepository.Exists(hash.String())
}

// Retrieve retrieves data by hash
func (app *repository) Retrieve(hsh hash.Hash) ([]byte, error) {
	return app.fileRepository.Retrieve(hsh.String())
}
