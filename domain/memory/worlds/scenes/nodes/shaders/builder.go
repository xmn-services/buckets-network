package shaders

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders/shader"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	shaders        []shader.Shader
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		hash:           nil,
		withoutHash:    false,
		shaders:        nil,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithoutHash flags the builder as without hash
func (app *builder) WithoutHash() Builder {
	app.withoutHash = true
	return app
}

// WithShaders add the shaders to the builder
func (app *builder) WithShaders(shaders []shader.Shader) Builder {
	app.shaders = shaders
	return app
}

// CreatedOn adds the creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn adds the lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds the new Shaders instance
func (app *builder) Now() (Shaders, error) {
	if app.shaders == nil {
		app.shaders = []shader.Shader{}
	}

	if app.withoutHash {
		data := [][]byte{
			[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
		}

		for _, oneShader := range app.shaders {
			data = append(data, oneShader.Hash().Bytes())
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Shaders instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createShaders(mutable, app.shaders), nil
}
