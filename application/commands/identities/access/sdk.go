package access

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/access/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder buckets.Builder,
) Builder {
	hashAdapter := hash.NewAdapter()
	accessBuilder := access.NewBuilder()
	return createBuilder(
		hashAdapter,
		accessBuilder,
		identityRepository,
		identityService,
		bucketAppBuilder,
	)
}

// Builder represents an access application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents an access application
type Application interface {
	Add(bucketHashStr string, privKey encryption.PrivateKey) error
	Retrieve(bucketHashStr string) (access.Access, error)
	Delete(bucketHashStr string) error
	Bucket(bucketHashStr string) (buckets.Application, error)
}
