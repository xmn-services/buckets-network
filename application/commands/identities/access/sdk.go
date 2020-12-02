package access

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/access/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
)

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
	Add(bucketHashStr string, pubKey public.Key) error
	Retrieve(bucketHashStr string) access.Access
	Delete(bucketHashStr string) error
	Bucket(bucketHashStr string) (buckets.Application, error)
}
