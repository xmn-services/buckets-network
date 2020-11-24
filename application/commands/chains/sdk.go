package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

// NewApplication creates a new application instance
func NewApplication(
	chainRepository chains.Repository,
) Application {
	return createApplication(chainRepository)
}

// Application represents the chain application
type Application interface {
	Retrieve() (chains.Chain, error)
	RetrieveAtIndex(index uint) (chains.Chain, error)
}
