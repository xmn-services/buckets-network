package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/xmn-services/buckets-network/libs/cryptography/encryption"
)

type encryptedFileDiskServiceBuilder struct {
	basePath string
	password string
}

func createEncryptedFileDiskServiceBuilder() EncryptedFileDiskServiceBuilder {
	out := encryptedFileDiskServiceBuilder{
		basePath: "",
		password: "",
	}

	return &out
}

// Create initializes the builder
func (app *encryptedFileDiskServiceBuilder) Create() EncryptedFileDiskServiceBuilder {
	return createEncryptedFileDiskServiceBuilder()
}

// WithBasePath adds a basePath to the builder
func (app *encryptedFileDiskServiceBuilder) WithBasePath(basePath string) EncryptedFileDiskServiceBuilder {
	app.basePath = basePath
	return app
}

// WithPassword adds a password to the builder
func (app *encryptedFileDiskServiceBuilder) WithPassword(password string) EncryptedFileDiskServiceBuilder {
	app.password = password
	return app
}

// Now builds a new Service instance
func (app *encryptedFileDiskServiceBuilder) Now() (Service, error) {
	if app.basePath == "" {
		return nil, errors.New("the basePath is mandatory in order to build a Service instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build a Service instance")
	}

	if st, err := os.Stat(app.basePath); !os.IsNotExist(err) {
		if !st.IsDir() {
			str := fmt.Sprintf("the basePath (%s) must be a directory", app.basePath)
			return nil, errors.New(str)
		}
	}

	encryption := encryption.NewEncryption(app.password)
	service := createFileDiskService(app.basePath)
	return createEncryptedFileDiskService(encryption, service), nil
}
