package genesis

// Application represents the genesis application
type Application interface {
	Current() Current
}

// Current represents the current application
type Current interface {
	Init(
		blockDifficultyBase uint,
		blockDifficultyIncreasePerTrx float64,
		linkDifficulty uint,
	) error
}
