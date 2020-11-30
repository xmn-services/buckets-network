package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	identityRepository identities.Repository
	identityService    identities.Service
	fileService        file.Service
	name               string
	password           string
	seed               string
}

func createApplication(
	hashAdapter hash.Adapter,
	identityRepository identities.Repository,
	identityService identities.Service,
	fileService file.Service,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		identityRepository: identityRepository,
		identityService:    identityService,
		fileService:        fileService,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Save saves a file
func (app *application) Save(file file.File) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// save the file:
	err = app.fileService.Save(file)
	if err != nil {
		return err
	}

	// add the file to the identity:
	err = identity.Wallet().Storage().Stored().Add(file.File().Hash())
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Delete deletes a file
func (app *application) Delete(fileHashStr string) error {
	fileHash, err := app.hashAdapter.FromString(fileHashStr)
	if err != nil {
		return err
	}

	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// delete the file:
	err = app.fileService.Delete(*fileHash)
	if err != nil {
		return err
	}

	// delete the file from the identity:
	err = identity.Wallet().Storage().Stored().Delete(*fileHash)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}
