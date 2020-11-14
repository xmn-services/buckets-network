package file

import (
	"errors"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents"
)

type builder struct {
	contentsBuilder contents.Builder
	file            files.File
	contents        [][]byte
}

func createBuilder(
	contentsBuilder contents.Builder,
) Builder {
	out := builder{
		contentsBuilder: contentsBuilder,
		file:            nil,
		contents:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.contentsBuilder)
}

// WithFile adds a file to the builder
func (app *builder) WithFile(file files.File) Builder {
	app.file = file
	return app
}

// WithContents adds contents to the builder
func (app *builder) WithContents(contents [][]byte) Builder {
	app.contents = contents
	return app
}

// Now builds a new File instance
func (app *builder) Now() (File, error) {
	if app.file == nil {
		return nil, errors.New("the File is mandatory in order to build a stored File instance")
	}

	contentsBuilder := app.contentsBuilder.Create()
	if app.contents != nil && len(app.contents) > 0 {
		contentsBuilder.WithContents(app.contents)
	}

	contents, err := contentsBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createFile(app.file, contents), nil
}
