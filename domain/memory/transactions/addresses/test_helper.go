package addresses

import "github.com/xmn-services/buckets-network/libs/file"

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() (Repository, Service) {
	fileRepositoryService := file.CreateRepositoryServiceForTests()
	repository := NewRepository(fileRepositoryService)
	service := NewService(fileRepositoryService)
	return repository, service
}
