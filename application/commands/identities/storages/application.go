package storages

import (
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	identityRepository identities.Repository
	bucketRepository   buckets.Repository
	contentService     contents.Service
	name               string
	password           string
	seed               string
}

func createApplication(
	hashAdapter hash.Adapter,
	identityRepository identities.Repository,
	bucketRepository buckets.Repository,
	contentService contents.Service,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		identityRepository: identityRepository,
		bucketRepository:   bucketRepository,
		contentService:     contentService,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Save saves a a chunk in bucket
func (app *application) Save(bucketHashStr string, chunk []byte) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// create an hash from the string:
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	// verify if the bucket exists for the authenticated identity:
	if identity.Wallet().Storage().Stored().Exists(*bucketHash) {
		str := fmt.Sprintf("the bucket (hash: %s) does not exists", bucketHash.String())
		return errors.New(str)
	}

	// retrieve the bucket:
	bucket, err := app.bucketRepository.Retrieve(*bucketHash)
	if err != nil {
		return err
	}

	// save the content:
	return app.contentService.Save(bucket, chunk)
}

// Delete deletes a hash from bucket
func (app *application) Delete(bucketHashStr string, chunkHashStr string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// create the bucket hash from the string:
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	// create the chunk hash from the string:
	chunkHash, err := app.hashAdapter.FromString(chunkHashStr)
	if err != nil {
		return err
	}

	// verify if the bucket exists for the authenticated identity:
	if identity.Wallet().Storage().Stored().Exists(*bucketHash) {
		str := fmt.Sprintf(bucketDoesNotExistsErr, bucketHash.String())
		return errors.New(str)
	}

	// retrieve the bucket:
	bucket, err := app.bucketRepository.Retrieve(*bucketHash)
	if err != nil {
		return err
	}

	// save the content:
	return app.contentService.Delete(bucket, *chunkHash)
}

// DeleteAll deletes all chunks from bucket
func (app *application) DeleteAll(bucketHashStr string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// create the bucket hash from the string:
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	// verify if the bucket exists for the authenticated identity:
	if identity.Wallet().Storage().Stored().Exists(*bucketHash) {
		str := fmt.Sprintf(bucketDoesNotExistsErr, bucketHash.String())
		return errors.New(str)
	}

	// retrieve the bucket:
	bucket, err := app.bucketRepository.Retrieve(*bucketHash)
	if err != nil {
		return err
	}

	// save the content:
	return app.contentService.DeleteAll(bucket)
}
