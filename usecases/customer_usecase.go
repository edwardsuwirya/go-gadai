package usecases

import (
	"enigmacamp.com/gosql/models"
	"enigmacamp.com/gosql/repositories"
	repo "enigmacamp.com/gosql/repositories"
	"mime/multipart"
)

type ICustomerUseCase interface {
	FindCustomerByFirstName(name string) ([]models.Customer, error)
	GetTotalCustomer() (int, error)
	RegisterNewCustomer(customer models.Customer) (*models.Customer, error)
	RegisterBulkCustomer(customer []models.Customer) ([]*models.Customer, error)
	UploadAvatar(fileName string, file multipart.File) error
}

type CustomerUseCase struct {
	repo     repo.ICustomerRepository
	fileRepo repo.IFileRepository
}

func NewCustomerUseCase(customerRepo repositories.ICustomerRepository, fileRepo repositories.IFileRepository) ICustomerUseCase {
	return &CustomerUseCase{
		repo:     customerRepo,
		fileRepo: fileRepo,
	}
}

func (c *CustomerUseCase) FindCustomerByFirstName(name string) ([]models.Customer, error) {
	return c.repo.FindAllByNameLike(name)
}

func (c *CustomerUseCase) GetTotalCustomer() (int, error) {
	return c.repo.Count()
}

func (c *CustomerUseCase) RegisterNewCustomer(customer models.Customer) (*models.Customer, error) {
	return c.repo.Insert(customer)
}

func (c *CustomerUseCase) UploadAvatar(fileName string, file multipart.File) error {
	return c.fileRepo.Save(file, fileName)
}

func (c *CustomerUseCase) RegisterBulkCustomer(newCustomers []models.Customer) ([]*models.Customer, error) {
	return c.repo.InsertBulk(newCustomers)
}
