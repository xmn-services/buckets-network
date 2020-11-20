package file

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents"
	transfer_data "github.com/xmn-services/buckets-network/domain/transfers/file"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewService creates a new service instance
func NewService(
	repository Repository,
	fileService files.Service,
	trDataService transfer_data.Service,
) Service {
	return createService(repository, fileService, trDataService)
}

// NewRepository creates a new repository instance
func NewRepository(
	fileRepository files.Repository,
	trDataRepository transfer_data.Repository,
) Repository {
	builder := NewBuilder()
	return createRepository(fileRepository, trDataRepository, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	contentsBuilder := contents.NewBuilder()
	return createBuilder(contentsBuilder)
}

// Builder represents a file builder
type Builder interface {
	Create() Builder
	WithFile(file files.File) Builder
	WithContents(contents [][]byte) Builder
	Now() (File, error)
}

// File represents a stored file
type File interface {
	File() files.File
	Contents() contents.Contents
}

// Repository represents a file repositroy
type Repository interface {
	Retrieve(fileHash hash.Hash) (File, error)
}

// Service represents a file service
type Service interface {
	Save(file File) error
	Delete(hash hash.Hash) error
}
