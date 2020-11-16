package genesis

// Application represents the genesis application
type Application interface {
	Init(baseDifficulty uint, increasePerBucket float64, linkDifficulty uint) error
}
