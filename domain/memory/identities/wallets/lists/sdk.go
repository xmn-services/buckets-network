package lists

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter returns the lists adapter
type Adapter interface {
	ToJSON(lists Lists) *JSONLists
	ToLists(ins *JSONLists) (Lists, error)
}

// Factory represents a lists factory
type Factory interface {
	Create() (Lists, error)
}

// Builder represents an lists builder
type Builder interface {
	Create() Builder
	WithList(lst []list.List) Builder
	Now() (Lists, error)
}

// Lists represents lists
type Lists interface {
	All() []list.List
	Fetch(listHash hash.Hash) (list.List, error)
	Add(list list.List) error
	Delete(listHash hash.Hash) error
	Update(list list.List) error
}
