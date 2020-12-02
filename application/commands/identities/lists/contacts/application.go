package contacts

import (
	application_contact_bucket "github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	identityRepository identities.Repository
	identityService    identities.Service
	bucketAppBuilder   application_contact_bucket.Builder
	name               string
	password           string
	seed               string
	listHash           hash.Hash
}

func createApplication(
	hashAdapter hash.Adapter,
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder application_contact_bucket.Builder,
	name string,
	password string,
	seed string,
	listHash hash.Hash,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		identityRepository: identityRepository,
		identityService:    identityService,
		bucketAppBuilder:   bucketAppBuilder,
		name:               name,
		password:           password,
		seed:               seed,
		listHash:           listHash,
	}

	return &out
}

// RetrieveAll retrieves the contacts of the list
func (app *application) RetrieveAll() (contacts.Contacts, error) {
	_, list, err := app.retrieveList()
	if err != nil {
		return nil, err
	}

	return list.Contacts(), nil
}

// Retrieve retrieves a contact by hash
func (app *application) Retrieve(contactHashStr string) (contact.Contact, error) {
	_, _, contact, err := app.retrieve(contactHashStr)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

// Update updates a contact
func (app *application) Update(contactHashStr string, update Update) error {
	identity, list, contact, err := app.retrieve(contactHashStr)
	if err != nil {
		return err
	}

	if update.HasKey() {
		contact.SetKey(update.Key())
	}

	if update.HasName() {
		contact.SetName(update.Name())
	}

	if update.HasDescription() {
		contact.SetDescription(update.Description())
	}

	return app.updateContact(identity, list, contact)
}

// Delete deletes a contact
func (app *application) Delete(contactHashStr string) error {
	contactHash, err := app.hashAdapter.FromString(contactHashStr)
	if err != nil {
		return err
	}

	identity, list, err := app.retrieveList()
	if err != nil {
		return err
	}

	err = list.Contacts().Delete(*contactHash)
	if err != nil {
		return err
	}

	return app.updateList(identity, list)
}

// Bucket creates a bucket contact application
func (app *application) Bucket(contactHashStr string) (application_contact_bucket.Application, error) {
	return app.bucketAppBuilder.Create().
		WithName(app.name).
		WithPassword(app.password).
		WithSeed(app.seed).
		WithList(app.listHash.String()).
		WithContact(contactHashStr).
		Now()
}

func (app *application) retrieve(contactHashStr string) (identities.Identity, list.List, contact.Contact, error) {
	contactHash, err := app.hashAdapter.FromString(contactHashStr)
	if err != nil {
		return nil, nil, nil, err
	}

	identity, list, err := app.retrieveList()
	if err != nil {
		return nil, nil, nil, err
	}

	contact, err := list.Contacts().Fetch(*contactHash)
	if err != nil {
		return nil, nil, nil, err
	}

	return identity, list, contact, nil
}

func (app *application) retrieveList() (identities.Identity, list.List, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, nil, err
	}

	// fetch the list from the identity:
	list, err := identity.Wallet().Lists().Fetch(app.listHash)
	if err != nil {
		return nil, nil, err
	}

	return identity, list, nil
}

func (app *application) updateList(identity identities.Identity, list list.List) error {
	// update the list in the identity:
	err := identity.Wallet().Lists().Update(list)
	if err != nil {
		return err
	}

	// save the identity:
	return app.identityService.Update(identity, app.password, app.password)
}

func (app *application) updateContact(identity identities.Identity, list list.List, cnt contact.Contact) error {
	// update the contact in the list:
	err := list.Contacts().Update(cnt)
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
