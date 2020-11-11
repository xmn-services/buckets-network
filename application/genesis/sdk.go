package genesis

// Application represents the genesis application
type Application interface {
	Init(
		blockDifficultyBase uint,
		blockDifficultyIncreasePerTrx float64,
		linkDifficulty uint,
	) error
}
