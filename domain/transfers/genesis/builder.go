package genesis

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	immutableBuilder        entities.ImmutableBuilder
	hash                    *hash.Hash
	blockDiffBase           uint
	blockDiffIncreasePerBucket float64
	linkDiff                uint
	createdOn               *time.Time
}

func createBuilder(
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		immutableBuilder:        immutableBuilder,
		hash:                    nil,
		blockDiffBase:           0,
		blockDiffIncreasePerBucket: 0,
		linkDiff:                0,
		createdOn:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.immutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithBlockDifficultyBase adds a block difficulty base to the builder
func (app *builder) WithBlockDifficultyBase(blockDiffBase uint) Builder {
	app.blockDiffBase = blockDiffBase
	return app
}

// WithBlockDifficultyIncreasePerBucket adds a block difficulty increase per bucket to the builder
func (app *builder) WithBlockDifficultyIncreasePerBucket(blockDiffIncreasePerBucket float64) Builder {
	app.blockDiffIncreasePerBucket = blockDiffIncreasePerBucket
	return app
}

// WithLinkDifficulty adds a link difficulty increase per bucket to the builder
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
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Genesis instance")
	}

	if app.blockDiffBase <= 0 {
		return nil, errors.New("the block difficulty base must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.blockDiffIncreasePerBucket <= 0 {
		return nil, errors.New("the block difficulty increasePerBucket must be greater than zero (0) in order to build a Genesis instance")
	}

	if app.linkDiff <= 0 {
		return nil, errors.New("the link difficulty must be greater than zero (0) in order to build a Genesis instance")
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createGenesis(
		immutable,
		app.blockDiffBase,
		app.blockDiffIncreasePerBucket,
		app.linkDiff,
	), nil
}
