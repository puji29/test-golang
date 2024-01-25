package usecase

import (
	"simple_payments/entity"
	"simple_payments/repository"
)

type MerchantUseCase interface {
	RegisterMer(name, password string) error
	FindByName(name string) error
}

type merchantUseCase struct {
	merchantRepo repository.MerchantRepository
}

// FindByName implements MerchantUseCase.
func (m *merchantUseCase) FindByName(name string) error {
	return m.FindByName(name)
}

// RegisterMer implements MerchantUseCase.
func (m *merchantUseCase) RegisterMer(name, password string) error {

	merchant := &entity.Merchant{
		Name:     name,
		Password: password,
		Balance:  0.0,
	}

	err := m.merchantRepo.RegisterMerchant(merchant)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantUseCase(merchantRepo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{merchantRepo: merchantRepo}
}
