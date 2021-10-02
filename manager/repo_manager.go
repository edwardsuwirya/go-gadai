package manager

import "enigmacamp.com/gosql/repositories"

type RepoManager interface {
	CustomerRepo() repositories.ICustomerRepository
}
type repoManager struct {
	db *repositories.DbSessionFactory
}

func (rm *repoManager) CustomerRepo() repositories.ICustomerRepository {
	return repositories.NewCustomerRepository(rm.db)
}

func NewRepoManager(sf *repositories.DbSessionFactory) RepoManager {
	return &repoManager{sf}
}
