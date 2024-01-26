package usecase

import (
	"simple_payments/entity"
	"simple_payments/repository"
)

type MerchantUseCase interface {
	RegisterMer(name, password string) error
	FindName(name string) (*entity.Merchant, error)
}

type merchantUseCase struct {
	merchantRepo repository.MerchantRepository
}

// FindByName implements MerchantUseCase.
func (u *merchantUseCase) FindName(name string) (*entity.Merchant, error) {
	return u.merchantRepo.FindByName(name)
}

// RegisterMer implements MerchantUseCase.
func (u *merchantUseCase) RegisterMer(name, password string) error {

	merchant := &entity.Merchant{
		Name:     name,
		Password: password,
		Balance:  0.0,
	}

	err := u.merchantRepo.RegisterMerchant(merchant)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantUseCase(merchantRepo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{merchantRepo: merchantRepo}
}
