package file

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
