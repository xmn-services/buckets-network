package transactions

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/transactions/addresses"
	transfer_transaction "github.com/xmn-services/buckets-network/domain/transfers/transactions"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	addressService addresses.Service,
	trService transfer_transaction.Service,
) Service {
	adapter := NewAdapter()
	return createService(adapter, repository, addressService, trService)
}

// NewRepository creates a new repository instance
func NewRepository(
	addressRepository addresses.Repository,
	trRepository transfer_transaction.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(builder, addressRepository, trRepository)
}

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	trBuilder := transfer_transaction.NewBuilder()
	return createAdapter(trBuilder)
}

// NewBuilder returns a new transaction builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Adapter returns the transaction adapter
type Adapter interface {
	ToTransfer(trx Transaction) (transfer_transaction.Transaction, error)
	ToJSON(trx Transaction) *JSONTransaction
	ToTransaction(ins *JSONTransaction) (Transaction, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithAddress(address addresses.Address) Builder
	WithBucket(bucket hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	entities.Immutable
	Bucket() hash.Hash
	HasAddress() bool
	Address() addresses.Address
}

// Repository represents a transaction repository
type Repository interface {
	Retrieve(hash hash.Hash) (Transaction, error)
	RetrieveAll(hashes []hash.Hash) ([]Transaction, error)
}

// Service represents the transaction service
type Service interface {
	Save(trx Transaction) error
	SaveAll(trx []Transaction) error
}
