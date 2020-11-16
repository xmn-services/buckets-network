package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/xmn-services/buckets-network/libs/cryptography/encryption"
)

type encryptedFileDiskRepositoryBuilder struct {
	basePath string
	password string
}

func createEncryptedFileDiskRepositoryBuilder() EncryptedFileDiskRepositoryBuilder {
	out := encryptedFileDiskRepositoryBuilder{
		basePath: "",
		password: "",
	}

	return &out
}

// Create initializes the builder
func (app *encryptedFileDiskRepositoryBuilder) Create() EncryptedFileDiskRepositoryBuilder {
	return createEncryptedFileDiskRepositoryBuilder()
}

// WithBasePath adds a basePath to the builder
func (app *encryptedFileDiskRepositoryBuilder) WithBasePath(basePath string) EncryptedFileDiskRepositoryBuilder {
	app.basePath = basePath
	return app
}

// WithPassword adds a password to the builder
func (app *encryptedFileDiskRepositoryBuilder) WithPassword(password string) EncryptedFileDiskRepositoryBuilder {
	app.password = password
	return app
}

// Now builds a new Repository instance
func (app *encryptedFileDiskRepositoryBuilder) Now() (Repository, error) {
	if app.basePath == "" {
		return nil, errors.New("the basePath is mandatory in order to build a Repository instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build a Repository instance")
	}

	if st, err := os.Stat(app.basePath); !os.IsNotExist(err) {
		if !st.IsDir() {
			str := fmt.Sprintf("the basePath (%s) must be a directory", app.basePath)
			return nil, errors.New(str)
		}
	}

	encryption := encryption.NewEncryption(app.password)
	repository := createFileDiskRepository(app.basePath)
	return createEncryptedFileDiskRepository(encryption, repository), nil
}
