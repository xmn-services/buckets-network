package file

// NewEncryptedFileDiskRepositoryBuilder creates a new encrypted file disk repository builder
func NewEncryptedFileDiskRepositoryBuilder() EncryptedFileDiskRepositoryBuilder {
	return createEncryptedFileDiskRepositoryBuilder()
}

// NewEncryptedFileDiskServiceBuilder creates a new encrypted file disk service builder
func NewEncryptedFileDiskServiceBuilder() EncryptedFileDiskServiceBuilder {
	return createEncryptedFileDiskServiceBuilder()
}

// NewFileDiskRepository creates a new repository that reads from files on disk
func NewFileDiskRepository(basePath string) Repository {
	return createFileDiskRepository(basePath)
}

// NewFileDiskService creates a new service that writes data on disk
func NewFileDiskService(basePath string) Service {
	return createFileDiskService(basePath)
}

// EncryptedFileDiskRepositoryBuilder represents an encrypted file fisk repository builder
type EncryptedFileDiskRepositoryBuilder interface {
	Create() EncryptedFileDiskRepositoryBuilder
	WithBasePath(basePath string) EncryptedFileDiskRepositoryBuilder
	WithPassword(password string) EncryptedFileDiskRepositoryBuilder
	Now() (Repository, error)
}

// EncryptedFileDiskServiceBuilder represents an encrypted file fisk service builder
type EncryptedFileDiskServiceBuilder interface {
	Create() EncryptedFileDiskServiceBuilder
	WithBasePath(basePath string) EncryptedFileDiskServiceBuilder
	WithPassword(password string) EncryptedFileDiskServiceBuilder
	Now() (Service, error)
}

// Repository represents a file repository
type Repository interface {
	Exists(relativePath string) bool
	Retrieve(relativePath string) ([]byte, error)
	RetrieveAll(relativePath string) ([]string, error)
}

// Service represents the file service
type Service interface {
	Save(relativePath string, content []byte) error
	Delete(relativePath string) error
	DeleteAll(relativePath []string) error
}
