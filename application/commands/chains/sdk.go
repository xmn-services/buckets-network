package chains

import (
	"github.com/xmn-services/buckets-network/application/commands/miners"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

// NewApplication creates a new application instance
func NewApplication(
	minerApp miners.Application,
	chainRepository chains.Repository,
	chainService chains.Service,
) Application {
	return createApplication(minerApp, chainRepository, chainService)
}

// Application represents the chain application
type Application interface {
	Init(
		miningValue uint8,
		baseDifficulty uint,
		increasePerBucket float64,
		linkDifficulty uint,
		rootAdditionalBuckets uint,
		headAdditionalBuckets uint,
	) error

	Retrieve() (chains.Chain, error)
	RetrieveAtIndex(index uint) (chains.Chain, error)
}
