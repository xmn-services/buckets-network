package chains

import (
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
)

// Application represents the chain application
type Application interface {
	Retrieve() (chains.Chain, error)
	RetrieveAtIndex(index uint) (chains.Chain, error)
	Upgrade(link mined_link.Link) error
}
