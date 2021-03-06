package file

import (
	"github.com/xmn-services/buckets-network/libs/cryptography/encryption"
)

type encryptedFileDiskRepository struct {
	encryption encryption.Encryption
	repository Repository
}

func createEncryptedFileDiskRepository(encryption encryption.Encryption, repository Repository) Repository {
	out := encryptedFileDiskRepository{
		encryption: encryption,
		repository: repository,
	}

	return &out
}

// Exists returns true if the file exists, false otherwise
func (app *encryptedFileDiskRepository) Exists(relativePath string) bool {
	return app.repository.Exists(relativePath)
}

// Retrieve retrieves data from file using its name
func (app *encryptedFileDiskRepository) Retrieve(relativePath string) ([]byte, error) {
	encrypted, err := app.repository.Retrieve(relativePath)
	if err != nil {
		return nil, err
	}

	return app.encryption.Decrypt(string(encrypted))
}

// RetrieveAll retrieves all files in a given directory
func (app *encryptedFileDiskRepository) RetrieveAll(relativePath string) ([]string, error) {
	return app.repository.RetrieveAll(relativePath)
}
