package layers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	layers         []layer.Layer
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
		layers:         nil,
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

// WithLayers add the layers to the builder
func (app *builder) WithLayers(layers []layer.Layer) Builder {
	app.layers = layers
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

// Now builds the new Layers instance
func (app *builder) Now() (Layers, error) {
	if app.layers == nil {
		app.layers = []layer.Layer{}
	}

	if app.withoutHash {
		data := [][]byte{
			[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
		}

		for _, oneLayer := range app.layers {
			data = append(data, oneLayer.Hash().Bytes())
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	totalAlpha := int32(0)
	for _, oneLayer := range app.layers {
		totalAlpha += int32(oneLayer.Alpha())
	}

	if totalAlpha >= maxAlpha {
		str := fmt.Sprintf("the layers cannot have a combined alpha greater than %d, %d provided", maxAlpha, totalAlpha)
		return nil, errors.New(str)
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Layers instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createLayers(mutable, app.layers), nil
}
