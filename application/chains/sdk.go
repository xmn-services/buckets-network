package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
)

// Application represents the chain application
type Application interface {
	Retrieve() (chains.Chain, error)
	RetrieveAtIndex(index uint) (chains.Chain, error)
	Update(updated chains.Chain) error
}
