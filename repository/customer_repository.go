package repository

import (
	"errors"
	"fmt"
	"simple_payments/entity"
	"sync"
)

var ErrCustomerNotFound = errors.New("customer not found")

type CustomerRepository interface {
	FindByUsername(username string) (*entity.Customer, error)
	Save(customer *entity.Customer) error
	Register(customer *entity.Customer) error
}

type InMemoryCustomerRepository struct {
	Customers map[string]entity.Customer
	mutex     sync.Mutex
}

func (repo *InMemoryCustomerRepository) FindByUsername(username string) (*entity.Customer, error) {
	customer, exists := repo.Customers[username]
	if !exists {
		return nil, ErrCustomerNotFound
	}
	return &customer, nil
}

func (repo *InMemoryCustomerRepository) Save(customer *entity.Customer) error {
	repo.Customers[customer.Username] = *customer
	return nil
}
func (repo *InMemoryCustomerRepository) Register(customer *entity.Customer) error {
	_, exists := repo.Customers[customer.Username]
	if exists {
		return fmt.Errorf("username already exists")
	}

	return repo.Save(customer)
}

func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{Customers: make(map[string]entity.Customer)}
}
