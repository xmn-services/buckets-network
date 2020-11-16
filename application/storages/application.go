package storages

import (
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter          hash.Adapter
	storedFileRepository stored_file.Repository
}

func createApplication(
	hashAdapter hash.Adapter,
	storedFileRepository stored_file.Repository,
) Application {
	out := application{
		hashAdapter:          hashAdapter,
		storedFileRepository: storedFileRepository,
	}

	return &out
}

// IsStored returns true if the file is stored, false otherwise
func (app *application) IsStored(fileHashStr string) bool {
	fileHash, err := app.hashAdapter.FromString(fileHashStr)
	if err != nil {
		return false
	}

	storedFile, err := app.storedFileRepository.Retrieve(*fileHash)
	if err != nil {
		return false
	}

	file := storedFile.File()
	return storedFile.Contents().NotStored(file) == uint(0)
}

// Retrieve retrieves a stored file, if exists
func (app *application) Retrieve(fileHashStr string) (stored_file.File, error) {
	fileHash, err := app.hashAdapter.FromString(fileHashStr)
	if err != nil {
		return nil, err
	}

	return app.storedFileRepository.Retrieve(*fileHash)
}
