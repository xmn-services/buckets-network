package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	identityRepository identities.Repository
	identityService    identities.Service
	fileService        file.Service
	name               string
	password           string
	seed               string
}

func createApplication(
	identityRepository identities.Repository,
	identityService identities.Service,
	fileService file.Service,
	name string,
	password string,
	seed string,
) Application {
	out := application{
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
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// save the file:
	err = app.fileService.Save(file)
	if err != nil {
		return err
	}

	// add the file to the identity:
	err = identity.Wallet().Storages().Stored().Add(file.File().Hash())
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Delete deletes a file
func (app *application) Delete(file hash.Hash) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
	if err != nil {
		return err
	}

	// delete the file:
	err = app.fileService.Delete(file)
	if err != nil {
		return err
	}

	// delete the file from the identity:
	err = identity.Wallet().Storages().Stored().Delete(file)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}
