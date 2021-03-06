package genesis

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter                hash.Adapter
	immutableBuilder           entities.ImmutableBuilder
	miningValue                uint8
	blockDiffBase              uint
	blockDiffIncreasePerBucket float64
	linkDiff                   uint
	createdOn                  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:                hashAdapter,
		immutableBuilder:           immutableBuilder,
		miningValue:                uint8(10),
		blockDiffBase:              0,
		blockDiffIncreasePerBucket: 0.0,
		linkDiff:                   0,
		createdOn:                  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithMiningValue adds a mining value to the builder
func (app *builder) WithMiningValue(miningValue uint8) Builder {
	app.miningValue = miningValue
	return app
}

// WithBlockDifficultyBase adds a block difficulty base to the builder
func (app *builder) WithBlockDifficultyBase(blockDiffBase uint) Builder {
	app.blockDiffBase = blockDiffBase
	return app
}

// WithBlockDifficultyIncreasePerBucket adds a block difficulty increasePerBucket to the builder
func (app *builder) WithBlockDifficultyIncreasePerBucket(blockDiffIncreasePerBucket float64) Builder {
	app.blockDiffIncreasePerBucket = blockDiffIncreasePerBucket
	return app
}

// WithLinkDifficulty adds a link difficulty to the builder
func (app *builder) WithLinkDifficulty(linkDiff uint) Builder {
	app.linkDiff = linkDiff
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.blockDiffBase <= 0 {
		return nil, errors.New("the block difficulty base must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.blockDiffIncreasePerBucket <= 0.0 {
		return nil, errors.New("the block difficulty increasePerBucket must be greater than zero (0.0) in order to build a Genesis instance")
	}

	if app.linkDiff <= 0 {
		return nil, errors.New("the link difficulty must be greater than zero (0.0) in order to build a Genesis instance")
	}

	if app.miningValue > 9 {
		return nil, errors.New("the miningValue must be a number between 0 and 9")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.blockDiffBase))),
		[]byte(strconv.FormatFloat(app.blockDiffIncreasePerBucket, 'f', 12, 64)),
		[]byte(strconv.Itoa(int(app.linkDiff))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	block := createBlock(app.blockDiffBase, app.blockDiffIncreasePerBucket)
	diff := createDifficulty(block, app.linkDiff)
	return createGenesis(immutable, app.miningValue, diff), nil
}
