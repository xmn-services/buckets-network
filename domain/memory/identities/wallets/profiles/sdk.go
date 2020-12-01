package profiles

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists"
)

// Builder represents a profile builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithLists(lists lists.Lists) Builder
	WithAccess(access accesses.Accesses) Builder
	Now() (Profile, error)
}

// Profile represents a profile
type Profile interface {
	Name() string
	Description() string
	List() lists.Lists
	Access() accesses.Accesses
}
