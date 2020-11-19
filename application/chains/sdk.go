package chains

import (
	"github.com/xmn-services/buckets-network/application/miners"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

// NewApplication creates a new application instance
func NewApplication(
	minerApp miners.Application,
	chainRepository chains.Repository,
	chainService chains.Service,
) Application {
	chainBuilder := chains.NewBuilder()
	return createApplication(minerApp, chainRepository, chainService, chainBuilder)
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
	Update(newMinedLink mined_link.Link) error
}
