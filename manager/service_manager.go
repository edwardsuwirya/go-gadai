package manager

import (
	"enigmacamp.com/gosql/repositories"
	"enigmacamp.com/gosql/usecases"
)

type ServiceManager interface {
	CustomerUseCase() usecases.ICustomerUseCase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) CustomerUseCase() usecases.ICustomerUseCase {
	return usecases.NewCustomerUseCase(sm.repo.CustomerRepo(), sm.repo.FileRepo())
}

func NewServiceManger(sf *repositories.DbSessionFactory, filePath string) ServiceManager {
	return &serviceManager{repo: NewRepoManager(sf, filePath)}
}
