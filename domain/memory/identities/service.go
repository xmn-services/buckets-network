package identities

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	hashAdapter        hash.Adapter
	fileServiceBuilder file.EncryptedFileDiskServiceBuilder
	repository         Repository
	basePath           string
	extension          string
}

func createService(
	hashAdapter hash.Adapter,
	fileServiceBuilder file.EncryptedFileDiskServiceBuilder,
	repository Repository,
	basePath string,
	extension string,
) Service {
	out := service{
		hashAdapter:        hashAdapter,
		fileServiceBuilder: fileServiceBuilder,
		repository:         repository,
		basePath:           basePath,
		extension:          extension,
	}

	return &out
}

// Insert inserts a new identity
func (app *service) Insert(identity Identity, password string) error {
	_, err := app.repository.Retrieve(identity.Name(), password, identity.Seed())
	if err == nil {
		str := fmt.Sprintf("the identity (name: %s) already exists", identity.Name())
		return errors.New(str)
	}

	return app.save(identity, password)
}

// Update updates an existing identity
func (app *service) Update(identity Identity, password string, newPassword string) error {
	_, err := app.repository.Retrieve(identity.Name(), password, identity.Seed())
	if err != nil {
		str := fmt.Sprintf("the identity (name: %s) does not exists and therefore cannot be updated", identity.Name())
		return errors.New(str)
	}

	return app.save(identity, password)
}

// Delete deletes an existing identity
func (app *service) Delete(identity Identity, password string) error {
	_, err := app.repository.Retrieve(identity.Name(), password, identity.Seed())
	if err != nil {
		str := fmt.Sprintf("the identity (name: %s) does not exists and therefore cannot be deleted", identity.Name())
		return errors.New(str)
	}

	pass, err := makePassword(app.hashAdapter, identity.Seed(), password)
	if err != nil {
		return err
	}

	fileService, err := app.fileServiceBuilder.Create().WithBasePath(app.basePath).WithPassword(pass).Now()
	if err != nil {
		return err
	}

	filePath := makeFileName(identity.Name(), app.extension)
	return fileService.Delete(filePath)
}

func (app *service) save(identity Identity, password string) error {
	pass, err := makePassword(app.hashAdapter, identity.Seed(), password)
	if err != nil {
		return err
	}

	fileService, err := app.fileServiceBuilder.Create().WithBasePath(app.basePath).WithPassword(pass).Now()
	if err != nil {
		return err
	}

	js, err := json.Marshal(identity)
	if err != nil {
		return err
	}

	filePath := makeFileName(identity.Name(), app.extension)
	return fileService.Save(filePath, js)
}
