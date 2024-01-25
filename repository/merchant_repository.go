package repository

import (
	"fmt"
	"simple_payments/entity"
	"sync"
)

type MerchantRepository interface {
	FindByName(name string) (*entity.Merchant, error)
	SaveMerchant(merchant *entity.Merchant) error
	RegisterMerchant(merchant *entity.Merchant) error
}

type InMemoryMerchantRepository struct {
	Merchant map[string]entity.Merchant
	mutex    sync.Mutex
}

func (repo *InMemoryMerchantRepository) FindByName(name string) (*entity.Merchant, error) {
	merchant, exist := repo.Merchant[name]
	if !exist {
		return nil, fmt.Errorf("merchant not found")
	}
	return &merchant, nil
}

func (repo *InMemoryMerchantRepository) SaveMerchant(merchant *entity.Merchant) error {

	repo.Merchant[merchant.Name] = *merchant
	return nil
}
func (repo *InMemoryMerchantRepository) RegisterMerchant(merchant *entity.Merchant) error {
	_, exists := repo.Merchant[merchant.Name]
	if exists {
		return fmt.Errorf("merchant already exists")
	}

	return repo.SaveMerchant(merchant)
}

func NewInMemoryMerchantRepository() *InMemoryMerchantRepository {
	return &InMemoryMerchantRepository{Merchant: make(map[string]entity.Merchant)}
}
