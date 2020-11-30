package identities

import "github.com/xmn-services/buckets-network/domain/memory/identities"

type current struct {
	identityRepository identities.Repository
	identityService    identities.Service
	name               string
	password           string
	seed               string
}

func createCurrent(
	identityRepository identities.Repository,
	identityService identities.Service,
	name string,
	password string,
	seed string,
) Current {
	out := current{
		identityRepository: identityRepository,
		identityService:    identityService,
		name:               name,
		password:           password,
		seed:               seed,
	}

	return &out
}

// Update updates the identity
func (app *current) Update(update Update) error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	// retrieve the identity:
	newPassword := app.password
	if update.HasSeed() {
		uSeed := update.Seed()
		identity.SetSeed(uSeed)
	}

	if update.HasName() {
		uName := update.Name()
		identity.SetName(uName)
	}

	if update.HasRoot() {
		uRoot := update.Root()
		identity.SetRoot(uRoot)
	}

	if update.HasPassword() {
		newPassword = update.Password()
	}

	err = app.identityService.Update(
		identity,
		app.password,
		newPassword,
	)

	if err != nil {
		return err
	}

	app.password = newPassword
	return nil
}

// Retrieve retrieves the identity
func (app *current) Retrieve() (identities.Identity, error) {
	return app.identityRepository.Retrieve(app.name, app.seed, app.password)
}

// Delete deletes the identity
func (app *current) Delete() error {
	// retrieve the identity:
	identity, err := app.identityRepository.Retrieve(app.name, app.seed, app.password)
	if err != nil {
		return err
	}

	return app.identityService.Delete(identity, app.password)
}
