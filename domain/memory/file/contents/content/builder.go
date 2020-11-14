package content

import (
	"errors"

	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content []byte) Builder {
	app.content = content
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Content instance")
	}

	hsh, err := app.hashAdapter.FromBytes(app.content)
	if err != nil {
		return nil, err
	}

	return createContent(*hsh, app.content), nil
}
