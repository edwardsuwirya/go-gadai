package usecases

import (
	"enigmacamp.com/gosql/models"
	"enigmacamp.com/gosql/repositories"
	repo "enigmacamp.com/gosql/repositories"
)

type ICustomerUseCase interface {
	FindCustomerByFirstName(name string) ([]models.Customer, error)
	GetTotalCustomer() (int, error)
	RegisterNewCustomer(customer models.Customer) (*models.Customer, error)
	RegisterBulkCustomer(customer []models.Customer) error
}

type CustomerUseCase struct {
	repo repo.ICustomerRepository
}

func NewCustomerUseCase(sf *repositories.DbSessionFactory) ICustomerUseCase {
	return &CustomerUseCase{
		repo: repo.NewCustomerRepository(sf),
	}
}

func (c *CustomerUseCase) FindCustomerByFirstName(name string) ([]models.Customer, error) {
	return c.repo.FindAllByNameLike(name)
}

func (c *CustomerUseCase) GetTotalCustomer() (int, error) {
	return c.repo.Count()
}

func (c *CustomerUseCase) RegisterNewCustomer(customer models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (c *CustomerUseCase) RegisterBulkCustomer(customer []models.Customer) error {
	panic("implement me")
}
