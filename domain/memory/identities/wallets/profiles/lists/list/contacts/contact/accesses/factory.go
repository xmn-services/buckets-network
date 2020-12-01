package accesses

import "github.com/xmn-services/buckets-network/libs/hash"

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new accesses instance
func (app *factory) Create() Accesses {
	lst := []hash.Hash{}
	mp := map[string]hash.Hash{}
	return createAccesses(lst, mp)
}
