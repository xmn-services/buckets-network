package chains

import "errors"

type builder struct {
	miningValue           uint8
	baseDifficulty        uint
	increasePerBucket     float64
	linkDifficulty        uint
	rootAdditionalBuckets uint
	headAdditionalBuckets uint
}

func createBuilder() Builder {
	out := builder{
		miningValue:           0,
		baseDifficulty:        0,
		increasePerBucket:     0.0,
		linkDifficulty:        0,
		rootAdditionalBuckets: 0,
		headAdditionalBuckets: 0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMiningValue adds a mining value to the builder
func (app *builder) WithMiningValue(miningValue uint8) Builder {
	app.miningValue = miningValue
	return app
}

// WithBaseDifficulty adds a base difficulty to the builder
func (app *builder) WithBaseDifficulty(baseDifficulty uint) Builder {
	app.baseDifficulty = baseDifficulty
	return app
}

// WithIncreasePerBucket adds an increasePerBucket to the builder
func (app *builder) WithIncreasePerBucket(increasePerBucket float64) Builder {
	app.increasePerBucket = increasePerBucket
	return app
}

// WithLinkDifficulty adds a link difficulty to the builder
func (app *builder) WithLinkDifficulty(linkDifficulty uint) Builder {
	app.linkDifficulty = linkDifficulty
	return app
}

// WithRootAdditionalBuckets adds a rootAdditionalBuckets to the builder
func (app *builder) WithRootAdditionalBuckets(rootAdditionalBuckets uint) Builder {
	app.rootAdditionalBuckets = rootAdditionalBuckets
	return app
}

// WithHeadAdditionalBuckets adds an headAdditionalBuckets to the builder
func (app *builder) WithHeadAdditionalBuckets(headAdditionalBuckets uint) Builder {
	app.headAdditionalBuckets = headAdditionalBuckets
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.baseDifficulty <= 0 {
		return nil, errors.New("the baseDifficulty is mandatory in order to build a Chain instance")
	}

	if app.increasePerBucket <= 0.0 {
		return nil, errors.New("the increasePerBucket is mandatory in order to build a Chain instance")
	}

	if app.linkDifficulty <= 0 {
		return nil, errors.New("the linkDifficulty is mandatory in order to build a Chain instance")
	}

	if app.rootAdditionalBuckets <= 0 {
		return nil, errors.New("the rootAdditionalBuckets is mandatory in order to build a Chain instance")
	}

	if app.headAdditionalBuckets <= 0 {
		return nil, errors.New("the headAdditionalBuckets is mandatory in order to build a Chain instance")
	}

	return createChain(
		app.miningValue,
		app.baseDifficulty,
		app.increasePerBucket,
		app.linkDifficulty,
		app.rootAdditionalBuckets,
		app.headAdditionalBuckets,
	), nil
}
