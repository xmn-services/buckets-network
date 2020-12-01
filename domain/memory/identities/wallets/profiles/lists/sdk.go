package lists

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a lists factory
type Factory interface {
	Create() Lists
}

// Lists represents lists
type Lists interface {
	All() []list.List
	Add(list list.List) error
	Delete(listHash hash.Hash) error
}
