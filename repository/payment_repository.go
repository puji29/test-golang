package repository

import (
	"simple_payments/entity"
	"sync"
)

type TransactionHistoryRepository interface {
	SaveTransaction(transaction *entity.Transaction) error
	GetAll() ([]entity.Transaction, error)
}

type InMemoryTransactionRepository struct {
	TransactionHistory []entity.Transaction
	mutex              sync.Mutex
}

func (repo *InMemoryTransactionRepository) SaveTransaction(transaction *entity.Transaction) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.TransactionHistory = append(repo.TransactionHistory, *transaction)
	return nil
}

func (repo *InMemoryTransactionRepository) GetAll() ([]entity.Transaction, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.TransactionHistory, nil
}

func NewInMemoryPaymentRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		TransactionHistory: make([]entity.Transaction, 0),
	}
}
