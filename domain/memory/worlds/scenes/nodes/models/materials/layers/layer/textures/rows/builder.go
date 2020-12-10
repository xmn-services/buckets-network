package rows

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	rows           []pixels.Pixels
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
		rows:           nil,
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

// WithRows add the rows to the builder
func (app *builder) WithRows(rows []pixels.Pixels) Builder {
	app.rows = rows
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

// Now builds the new Rows instance
func (app *builder) Now() (Rows, error) {
	if app.rows == nil {
		app.rows = []pixels.Pixels{}
	}

	if len(app.rows) > 0 {
		amount := app.rows[0].Amount()
		for index, onePixels := range app.rows {
			pixelsAmount := onePixels.Amount()
			if amount == pixelsAmount {
				continue
			}

			str := fmt.Sprintf("the pixels were expected to contain %d pixels, %d provided at index: %d", amount, pixelsAmount, index)
			return nil, errors.New(str)
		}
	}

	if app.withoutHash {
		data := [][]byte{
			[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
		}

		for _, oneRow := range app.rows {
			data = append(data, oneRow.Hash().Bytes())
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Rows instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createRows(mutable, app.rows), nil
}
