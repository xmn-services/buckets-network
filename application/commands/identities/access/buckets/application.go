package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	identityRepository identities.Repository
	identityService    identities.Service
	bucketRepository   buckets.Repository
	contentService     contents.Service
	name               string
	password           string
	seed               string
	bucketHash         hash.Hash
}

func createApplication(
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketRepository buckets.Repository,
	contentService contents.Service,
	name string,
	password string,
	seed string,
	bucketHash hash.Hash,
) Application {
	out := application{
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketRepository:   bucketRepository,
		contentService:     contentService,
		name:               name,
		password:           password,
		seed:               seed,
		bucketHash:         bucketHash,
	}

	return &out
}

// Delete deletes the bucket
func (app *application) Delete() error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// delete the bucket:
	err = identity.Wallet().Accesses().Delete(app.bucketHash)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity, app.password, app.password)
}

// Retrieve retrieves the bucket
func (app *application) Retrieve() (buckets.Bucket, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	// retrieve the access:
	access, err := identity.Wallet().Accesses().Fetch(app.bucketHash)
	if err != nil {
		return nil, err
	}

	// retrieve the bucket:
	bucketHash := access.Bucket()
	return app.bucketRepository.Retrieve(bucketHash)
}

// Extract extracts the bucket to the path
func (app *application) Extract(absolutePath string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// retrieve the access:
	access, err := identity.Wallet().Accesses().Fetch(app.bucketHash)
	if err != nil {
		return err
	}

	// retrieve the bucket:
	bucketHash := access.Bucket()
	bucket, err := app.bucketRepository.Retrieve(bucketHash)
	if err != nil {
		return err
	}

	pubKey := access.Key()
	return app.contentService.Extract(bucket, pubKey, absolutePath)
}
