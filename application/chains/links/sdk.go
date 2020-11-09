package links

import (
	mined_link "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Application represents the link application
type Application interface {
	Retrieve(hash hash.Hash) (mined_link.Link, error)
}
