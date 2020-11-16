package identities

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets"
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(basePath string, extension string) Service {
	hashAdapter := hash.NewAdapter()
	fileServiceBuilder := file.NewEncryptedFileDiskServiceBuilder()
	repository := NewRepository(basePath, extension)
	return createService(hashAdapter, fileServiceBuilder, repository, basePath, extension)
}

// NewRepository creates a new repository instance
func NewRepository(basePath string, extension string) Repository {
	hashAdapter := hash.NewAdapter()
	fileRepositoryBuilder := file.NewEncryptedFileDiskRepositoryBuilder()
	return createRepository(hashAdapter, fileRepositoryBuilder, basePath, extension)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	walletFactory := wallets.NewFactory()
	return createBuilder(walletFactory)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithSeed(seed string) Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithWallet(wallet wallets.Wallet) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	Seed() string
	SetSeed(seed string)
	Name() string
	SetName(name string)
	Root() string
	SetRoot(root string)
	Wallet() wallets.Wallet
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(name string, password string, seed string) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Insert(identity Identity, password string) error
	Update(identity Identity, password string, newPassword string) error
	Delete(identity Identity, password string) error
}
