package files

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files/contents"
)

type builder struct {
	contents []contents.Content
}

func createBuilder() Builder {
	out := builder{
		contents: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContents add contents to the builder
func (app *builder) WithContents(contents []contents.Content) Builder {
	app.contents = contents
	return app
}

// Now builds a new Files instance
func (app *builder) Now() (Files, error) {
	if app.contents == nil {
		app.contents = []contents.Content{}
	}

	mp := map[string]contents.Content{}
	for _, oneContent := range app.contents {
		keyname := oneContent.Chunk().String()
		mp[keyname] = oneContent
	}

	return createFiles(app.contents, mp), nil
}
