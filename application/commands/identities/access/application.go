package access

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/access/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	accessBuilder      access.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	bucketAppBuilder   buckets.Builder
	name               string
	password           string
	seed               string
}

func createApplication(
	hashAdapter hash.Adapter,
	accessBuilder access.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder buckets.Builder,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		accessBuilder:      accessBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketAppBuilder:   bucketAppBuilder,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Add adds a bucket to the access
func (app *application) Add(bucketHashStr string, privKey encryption.PrivateKey) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	access, err := app.accessBuilder.Create().WithBucket(*bucketHash).WithKey(privKey).Now()
	if err != nil {
		return err
	}

	// add the access to the identity:
	err = identity.Wallet().Accesses().Add(access)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Retrieve retrieves an access by bucket hash
func (app *application) Retrieve(bucketHashStr string) (access.Access, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return nil, err
	}

	return identity.Wallet().Accesses().Fetch(*bucketHash)
}

// Delete deletes an access by bucket hash
func (app *application) Delete(bucketHashStr string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	// delete the identity:
	err = identity.Wallet().Accesses().Delete(*bucketHash)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

// Bucket creates a bucket application
func (app *application) Bucket(bucketHashStr string) (buckets.Application, error) {
	return app.bucketAppBuilder.Create().
		WithName(app.name).
		WithSeed(app.seed).
		WithPassword(app.password).
		WithBucket(bucketHashStr).
		Now()
}
