package contents

import (
	"github.com/xmn-services/buckets-network/libs/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service
func NewService(
	fileService file.Service,
) Service {
	hashAdapter := hash.NewAdapter()
	return createService(hashAdapter, fileService)
}

// NewRepository creates a new repository
func NewRepository(
	fileRepository file.Repository,
) Repository {
	return createRepository(fileRepository)
}

// Repository represents a content repository
type Repository interface {
	Retrieve(bucketHash hash.Hash, fileHash hash.Hash, chunkHash hash.Hash) ([]byte, error)
}

// Service represents a content service
type Service interface {
	Save(bucketHash hash.Hash, fileHash hash.Hash, data []byte) error
	Delete(bucketHash hash.Hash, fileHash hash.Hash, chunkHash hash.Hash) error
	DeleteFile(bucketHash hash.Hash, fileHash hash.Hash) error
	DeleteAll(bucketHash hash.Hash) error
}
