package pixels

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	pixels         []pixel.Pixel
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
		pixels:         nil,
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

// WithPixels add the pixels to the builder
func (app *builder) WithPixels(pixels []pixel.Pixel) Builder {
	app.pixels = pixels
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

// Now builds the new Pixels instance
func (app *builder) Now() (Pixels, error) {
	if app.pixels == nil {
		app.pixels = []pixel.Pixel{}
	}

	if app.withoutHash {
		data := [][]byte{
			[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
		}

		for _, onePixel := range app.pixels {
			data = append(data, []byte(onePixel.String()))
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Pixels instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createPixels(mutable, app.pixels), nil
}
