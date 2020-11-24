package chains

type chain struct {
	miningValue           uint8
	baseDifficulty        uint
	increasePerBucket     float64
	linkDifficulty        uint
	rootAdditionalBuckets uint
	headAdditionalBuckets uint
}

func createChain(
	miningValue uint8,
	baseDifficulty uint,
	increasePerBucket float64,
	linkDifficulty uint,
	rootAdditionalBuckets uint,
	headAdditionalBuckets uint,
) Chain {
	out := chain{
		miningValue:           miningValue,
		baseDifficulty:        baseDifficulty,
		increasePerBucket:     increasePerBucket,
		linkDifficulty:        linkDifficulty,
		rootAdditionalBuckets: rootAdditionalBuckets,
		headAdditionalBuckets: headAdditionalBuckets,
	}

	return &out
}

// MiningValue returns the mining value
func (obj *chain) MiningValue() uint8 {
	return obj.miningValue
}

// BaseDifficulty returns the base difficulty
func (obj *chain) BaseDifficulty() uint {
	return obj.baseDifficulty
}

// IncreasePerBucket returns the increase per bucket
func (obj *chain) IncreasePerBucket() float64 {
	return obj.increasePerBucket
}

// LinkDifficulty returns the link difficulty
func (obj *chain) LinkDifficulty() uint {
	return obj.linkDifficulty
}

// RootAdditionalBuckets returns the root additional buckets
func (obj *chain) RootAdditionalBuckets() uint {
	return obj.rootAdditionalBuckets
}

// HeadAdditionalBuckets returns the head additional buckets
func (obj *chain) HeadAdditionalBuckets() uint {
	return obj.headAdditionalBuckets
}
