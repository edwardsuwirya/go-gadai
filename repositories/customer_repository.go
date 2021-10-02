package repositories

import (
	"database/sql"
	"enigmacamp.com/gosql/logger"
	"enigmacamp.com/gosql/models"
	"errors"
	guuid "github.com/google/uuid"
)

type ICustomerRepository interface {
	Insert(customer models.Customer) (*models.Customer, error)
	InsertBulk(customers []models.Customer) ([]*models.Customer, error)
	FindOneById(id string) (*models.Customer, error)
	FindAllByNameLike(name string) ([]models.Customer, error)
	Count() (int, error)
}

type CustomerRepository struct {
	dbSession *DbSessionFactory
}

func NewCustomerRepository(sf *DbSessionFactory) ICustomerRepository {
	customerRepo := &CustomerRepository{
		dbSession: sf,
	}
	return customerRepo
}

func (c *CustomerRepository) Insert(customer models.Customer) (*models.Customer, error) {
	newCustomerId := guuid.New()
	customer.Id = newCustomerId.String()
	_, err := c.dbSession.Db.Exec("INSERT INTO m_customer values (?,?,?,?,?)", customer.Id, customer.FirstName, customer.LastName, customer.Address, customer.City)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return nil, err
	} else {
		return &customer, nil
	}
}

func (c *CustomerRepository) InsertBulk(newCustomers []models.Customer) (customers []*models.Customer, err error) {
	tx, err := c.dbSession.Db.Begin()
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return nil, err
	}
	defer func(tx *sql.Tx) {
		if e := recover(); e != nil {
			tx.Rollback()
			err = errors.New(e.(string))
			logger.Logger.Error().Msg(err.Error())
			customers = nil
		}
	}(tx)

	stmt, err := tx.Prepare("INSERT INTO m_customer values (?,?,?,?,?)")
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		e := stmt.Close()
		if e != nil {
			logger.Logger.Error().Msg(err.Error())
			err = e
			customers = nil
		}
	}(stmt)
	for _, customer := range newCustomers {
		newCustomerId := guuid.New()
		customer.Id = newCustomerId.String()
		_, err := stmt.Exec(customer.Id, customer.FirstName, customer.LastName, customer.Address, customer.City)
		panic("Test")
		if err != nil {
			logger.Logger.Error().Msg(err.Error())
			return nil, err
		} else {
			customers = append(customers, &customer)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return
}

func (c *CustomerRepository) FindOneById(id string) (*models.Customer, error) {
	panic("implement me")
}

func (c *CustomerRepository) FindAllByNameLike(name string) (customers []models.Customer, err error) {
	rows, err := c.dbSession.Db.Query("select id,first_name,last_name,address,city from m_customer where first_name like ?", name)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		e := rows.Close()
		if e != nil {
			err = e
			customers = nil
		}
	}(rows)
	for rows.Next() {
		res := models.Customer{}
		err := rows.Scan(&res.Id, &res.FirstName, &res.LastName, &res.Address, &res.City)
		if err != nil {
			logger.Logger.Error().Msg(err.Error())
			return nil, err
		}
		customers = append(customers, res)
	}
	return
}

func (c *CustomerRepository) Count() (int, error) {
	var totalRecord int
	err := c.dbSession.Db.QueryRow("select count(*) from m_customer").Scan(&totalRecord)
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return -1, err
	}
	return totalRecord, nil

}
