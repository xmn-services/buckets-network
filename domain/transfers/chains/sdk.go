package chains

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(fileService file.Service, fileName string, extName string) Service {
	adapter := NewAdapter()
	return createService(adapter, fileService, fileName, extName)
}

// NewRepository creates a new repository instance
func NewRepository(fileRepository file.Repository, fileName string, extName string) Repository {
	adapter := NewAdapter()
	return createRepository(adapter, fileRepository, fileName, extName)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(immutableBuilder)
}

// Adapter represents a chain adapter
type Adapter interface {
	ToChain(js []byte) (Chain, error)
	ToJSON(chain Chain) ([]byte, error)
}

// Builder represents the chain builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithGenesis(gen hash.Hash) Builder
	WithRoot(root hash.Hash) Builder
	WithHead(head hash.Hash) Builder
	WithTotal(total uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	entities.Immutable
	Genesis() hash.Hash
	Root() hash.Hash
	Head() hash.Hash
	Total() uint
}

// Repository represents the chain repository
type Repository interface {
	Retrieve() (Chain, error)
	RetrieveAtIndex(index uint) (Chain, error)
}

// Service represents the chain service
type Service interface {
	Save(chain Chain, index uint) error
}
