package files

import "github.com/xmn-services/buckets-network/libs/hash"

type builder struct {
	lst []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithFiles add files to the builder
func (app *builder) WithFiles(hashes []hash.Hash) Builder {
	app.lst = hashes
	return app
}

// Now builds a new Files instance
func (app *builder) Now() (Files, error) {
	if app.lst == nil {
		app.lst = []hash.Hash{}
	}

	mp := map[string]hash.Hash{}
	for _, oneHash := range app.lst {
		keyname := oneHash.String()
		mp[keyname] = oneHash
	}

	return createFiles(app.lst, mp), nil
}
