package accesses

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses/access"

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new accesses instance
func (app *factory) Create() Accesses {
	lst := []access.Access{}
	mp := map[string]access.Access{}
	return createAccesses(lst, mp)
}
