package chains

import "net/url"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an init chain adapter
type Adapter interface {
	URLValuesToChain(values url.Values) (Chain, error)
	ChainToURLValues(chain Chain) url.Values
}

// Builder represents the init builder
type Builder interface {
	Create() Builder
	WithMiningValue(miningValue uint8) Builder
	WithBaseDifficulty(baseDiff uint) Builder
	WithIncreasePerBucket(incrPerBucket float64) Builder
	WithLinkDifficulty(likDiff uint) Builder
	WithRootAdditionalBuckets(rootAddBuckets uint) Builder
	WithHeadAdditionalBuckets(headAddBuckets uint) Builder
	Now() (Chain, error)
}

// Chain represents the init Chain
type Chain interface {
	MiningValue() uint8
	BaseDifficulty() uint
	IncreasePerBucket() float64
	LinkDifficulty() uint
	RootAdditionalBuckets() uint
	HeadAdditionalBuckets() uint
}
