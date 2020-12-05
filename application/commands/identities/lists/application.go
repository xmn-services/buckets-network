package lists

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type application struct {
	hashAdapter        hash.Adapter
	listBuilder        list.Builder
	identityRepository identities.Repository
	identityService    identities.Service
	contactsAppBuilder contacts.Builder
	name               string
	password           string
	seed               string
}

func createApplication(
	hashAdapter hash.Adapter,
	listBuilder list.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
	contactsAppBuilder contacts.Builder,
	name string,
	password string,
	seed string,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		listBuilder:        listBuilder,
		identityRepository: identityRepository,
		identityService:    identityService,
		contactsAppBuilder: contactsAppBuilder,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// RetrieveAll retrieves the lists
func (app *application) RetrieveAll() (lists.Lists, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	return identity.Wallet().Lists(), nil
}

// Retrieve retrieves the list by hash
func (app *application) Retrieve(listHashStr string) (list.List, error) {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return nil, err
	}

	listHash, err := app.hashAdapter.FromString(listHashStr)
	if err != nil {
		return nil, err
	}

	return identity.Wallet().Lists().Fetch(*listHash)
}

// New creates a new list
func (app *application) New(name string, description string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	list, err := app.listBuilder.Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		return err
	}

	err = identity.Wallet().Lists().Add(list)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity, app.password, app.password)
}

// Update updates an existing list
func (app *application) Update(listHashStr string, update Update) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	listHash, err := app.hashAdapter.FromString(listHashStr)
	if err != nil {
		return err
	}

	list, err := identity.Wallet().Lists().Fetch(*listHash)
	if err != nil {
		return err
	}

	if update.HasName() {
		list.SetName(update.Name())
	}

	if update.HasDescription() {
		list.SetDescription(update.Description())
	}

	err = identity.Wallet().Lists().Update(list)
	if err != nil {
		return err
	}

	return app.identityService.Update(identity, app.password, app.password)
}

// Delete deletes an existing list
func (app *application) Delete(listHashStr string) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	listHash, err := app.hashAdapter.FromString(listHashStr)
	if err != nil {
		return err
	}

	return identity.Wallet().Lists().Delete(*listHash)
}

// Contacts creates a contacts application instance
func (app *application) Contacts(listHashStr string) (contacts.Application, error) {
	list, err := app.Retrieve(listHashStr)
	if err != nil {
		return nil, err
	}

	return app.contactsAppBuilder.Create().WithName(app.name).WithPassword(app.password).WithSeed(app.seed).WithList(list.Hash().String()).Now()
}
