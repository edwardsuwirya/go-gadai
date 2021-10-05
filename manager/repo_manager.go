package manager

import (
	"enigmacamp.com/gosql/repositories"
)

type RepoManager interface {
	CustomerRepo() repositories.ICustomerRepository
	FileRepo() repositories.IFileRepository
}
type repoManager struct {
	db       *repositories.DbSessionFactory
	filePath string
}

func (rm *repoManager) CustomerRepo() repositories.ICustomerRepository {
	return repositories.NewCustomerRepository(rm.db)
}
func (rm *repoManager) FileRepo() repositories.IFileRepository {
	return repositories.NewFileRepository(rm.filePath)
}

func NewRepoManager(sf *repositories.DbSessionFactory, filePath string) RepoManager {
	return &repoManager{sf, filePath}
}
