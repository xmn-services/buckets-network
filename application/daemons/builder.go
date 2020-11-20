package daemons

import (
	"errors"
	"time"

	"github.com/xmn-services/buckets-network/application/syncs"
)

type builder struct {
	syncBuilder               syncs.Builder
	name                      string
	password                  string
	seed                      string
	waitPeriod                *time.Duration
	additionalBucketsPerBlock uint
}

func createBuilder(
	syncBuilder syncs.Builder,
) Builder {
	out := builder{
		syncBuilder:               syncBuilder,
		name:                      "",
		password:                  "",
		seed:                      "",
		waitPeriod:                nil,
		additionalBucketsPerBlock: uint(0),
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.syncBuilder,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithWaitPeriod adds a wait period to the builder
func (app *builder) WithWaitPeriod(waitPeriod time.Duration) Builder {
	app.waitPeriod = &waitPeriod
	return app
}

// WithAdditionalBucketsPerBlock adds an additional buckets per block to the builder
func (app *builder) WithAdditionalBucketsPerBlock(additionalBucketsPerBlock uint) Builder {
	app.additionalBucketsPerBlock = additionalBucketsPerBlock
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.waitPeriod == nil {
		return nil, errors.New("the waitPeriod is mandatory in order to build an Application instance")
	}

	syncApp, err := app.syncBuilder.Create().
		WithName(app.name).
		WithPassword(app.password).
		WithSeed(app.seed).
		WithAdditionalBucketsPerBlock(app.additionalBucketsPerBlock).
		Now()

	if err != nil {
		return nil, err
	}

	return createApplication(
		syncApp,
		*app.waitPeriod,
	), nil
}
