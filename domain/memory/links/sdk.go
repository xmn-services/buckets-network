package links

import (
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	blockService mined_blocks.Service,
	trService transfer_link.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, blockService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	blockRepository mined_blocks.Repository,
	trRepository transfer_link.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, blockRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_link.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the link adapter
type Adapter interface {
	ToTransfer(link Link) (transfer_link.Link, error)
	ToJSON(link Link) *JSONLink
	ToLink(ins *JSONLink) (Link, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithPrevious(prev hash.Hash) Builder
	WithNext(next mined_blocks.Block) Builder
	WithIndex(index uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Link, error)
}

// Link represents a block link
type Link interface {
	entities.Immutable
	Previous() hash.Hash
	Next() mined_blocks.Block
	Index() uint
}

// Repository represents a link repository
type Repository interface {
	Retrieve(hash hash.Hash) (Link, error)
}

// Service represents the link service
type Service interface {
	Save(link Link) error
}
