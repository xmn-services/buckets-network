package storages

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	storedFileRepository stored_file.Repository
}

func createApplication(
	storedFileRepository stored_file.Repository,
) Application {
	out := application{
		storedFileRepository: storedFileRepository,
	}

	return &out
}

// IsStored returns true if the file is stored, false otherwise
func (app *application) IsStored(fileHash hash.Hash) bool {
	storedFile, err := app.storedFileRepository.Retrieve(fileHash)
	if err != nil {
		return false
	}

	file := storedFile.File()
	return storedFile.Contents().NotStored(file) == uint(0)
}

// Retrieve retrieves a stored file, if exists
func (app *application) Retrieve(fileHash hash.Hash) (stored_file.File, error) {
	return app.storedFileRepository.Retrieve(fileHash)
}
