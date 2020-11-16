package storages

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"

type builder struct {
	filesFactory files.Factory
	toDownload   files.Files
	stored       files.Files
}

func createBuilder(
	filesFactory files.Factory,
) Builder {
	out := builder{
		filesFactory: filesFactory,
		toDownload:   nil,
		stored:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.filesFactory)
}

// WithToDownload adds toDownload files to the builder
func (app *builder) WithToDownload(toDownload files.Files) Builder {
	app.toDownload = toDownload
	return app
}

// WithStored adds stored files to the builder
func (app *builder) WithStored(stored files.Files) Builder {
	app.stored = stored
	return app
}

// Now builds a new Storage instance
func (app *builder) Now() (Storage, error) {
	if app.toDownload == nil {
		toDownload, err := app.filesFactory.Create()
		if err != nil {
			return nil, err
		}

		app.toDownload = toDownload
	}

	if app.stored == nil {
		stored, err := app.filesFactory.Create()
		if err != nil {
			return nil, err
		}

		app.stored = stored
	}

	return createStorage(app.toDownload, app.stored), nil
}
