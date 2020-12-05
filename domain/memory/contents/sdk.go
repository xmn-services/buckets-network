package contents

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	transfer_content "github.com/xmn-services/buckets-network/domain/transfers/contents"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service
func NewService(repository Repository, trService transfer_content.Service) Service {
	hashAdapter := hash.NewAdapter()
	return createService(hashAdapter, repository, trService)
}

// NewRepository creates a new repository
func NewRepository(trRepository transfer_content.Repository) Repository {
	return createRepository(trRepository)
}

// Repository represents a content repository
type Repository interface {
	Retrieve(bucket buckets.Bucket, fileHash hash.Hash, chunkHash hash.Hash) ([]byte, error)
}

// Service represents a content service
type Service interface {
	Extract(bucket buckets.Bucket, decryptPrivKey encryption.PrivateKey, absolutePath string) error
	Save(bucket buckets.Bucket, data []byte) error
	Delete(bucket buckets.Bucket, chunkHash hash.Hash) error
	DeleteAll(bucket buckets.Bucket) error
}
