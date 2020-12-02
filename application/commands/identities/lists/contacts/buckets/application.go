package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	contentService     contents.Service
	bucketRepository   buckets.Repository
	bucketService      buckets.Service
	identityRepository identities.Repository
	identityService    identities.Service
	name               string
	password           string
	seed               string
	listHash           hash.Hash
	contactHash        hash.Hash
}

func createApplication(
	hashAdapter hash.Adapter,
	contentService contents.Service,
	bucketRepository buckets.Repository,
	bucketService buckets.Service,
	identityRepository identities.Repository,
	identityService identities.Service,
	name string,
	password string,
	seed string,
	listHash hash.Hash,
	contactHash hash.Hash,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		contentService:     contentService,
		bucketRepository:   bucketRepository,
		bucketService:      bucketService,
		identityRepository: identityRepository,
		identityService:    identityService,
		name:               name,
		password:           password,
		seed:               seed,
		listHash:           listHash,
		contactHash:        contactHash,
	}

	return &out
}

// Add adds a bucket to the contact
func (app *application) Add(absolutePath string) error {
	// retrieve the contact:
	identity, list, contact, err := app.retrieveContact()
	if err != nil {
		return err
	}

	// retrieve the bucket:
	pubKey := contact.Key()
	bucket, chunksContent, err := app.bucketRepository.RetrieveWithChunksContentFromPath(absolutePath, pubKey)
	if err != nil {
		return err
	}

	// save the bucket:
	err = app.bucketService.Save(bucket)
	if err != nil {
		return err
	}

	// save the content:
	for _, oneFileContents := range chunksContent {
		for _, oneChunkContent := range oneFileContents {
			err := app.contentService.Save(bucket, oneChunkContent)
			if err != nil {
				return err
			}
		}
	}

	// update the contact:
	return app.updateContact(identity, list, contact)
}

// Delete deletes a bucket by hash
func (app *application) Delete(bucketHashStr string) error {
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return err
	}

	// retrieve the contact:
	identity, list, contact, err := app.retrieveContact()
	if err != nil {
		return err
	}

	// delete the access for the contact for the bucket:
	err = contact.Access().Delete(*bucketHash)
	if err != nil {
		return err
	}

	// update the contact:
	return app.updateContact(identity, list, contact)
}

// Retrieve retrieves a bucket by hash
func (app *application) Retrieve(bucketHashStr string) (buckets.Bucket, error) {
	bucketHash, err := app.hashAdapter.FromString(bucketHashStr)
	if err != nil {
		return nil, err
	}

	// retrieve the contact:
	_, _, contact, err := app.retrieveContact()
	if err != nil {
		return nil, err
	}

	// verify if the contact has access to the bucket:
	_, err = contact.Access().Fetch(*bucketHash)
	if err != nil {
		return nil, err
	}

	// retrieve the bucket:
	return app.bucketRepository.Retrieve(*bucketHash)
}

// RetrieveAll retrieves the contact's buckets
func (app *application) RetrieveAll() ([]buckets.Bucket, error) {
	// retrieve the contact:
	_, _, contact, err := app.retrieveContact()
	if err != nil {
		return nil, err
	}

	// retrieve the bucket hashes that the contact has access to:
	bucketHashes := contact.Access().All()

	// retrieve the buckets:
	buckets := []buckets.Bucket{}
	for _, oneBucketHash := range bucketHashes {
		bucket, err := app.bucketRepository.Retrieve(oneBucketHash)
		if err != nil {
			return nil, err
		}

		buckets = append(buckets, bucket)
	}

	// return the buckets:
	return buckets, nil
}

func (app *application) updateContact(identity identities.Identity, list list.List, contact contact.Contact) error {
	// update the contact in the list:
	err := list.Contacts().Update(contact)
	if err != nil {
		return err
	}

	// update the list in the identity:
	err = identity.Wallet().Lists().Update(list)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

func (app *application) retrieveContact() (identities.Identity, list.List, contact.Contact, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, nil, nil, err
	}

	// fetch the list from the identity:
	list, err := identity.Wallet().Lists().Fetch(app.listHash)
	if err != nil {
		return nil, nil, nil, err
	}

	// retrieve the contact from the list:
	contact, err := list.Contacts().Fetch(app.contactHash)
	if err != nil {
		return nil, nil, nil, err
	}

	return identity, list, contact, nil
}
