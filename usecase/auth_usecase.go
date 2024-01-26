package usecase

import (
	"fmt"
	"log"
	"simple_payments/entity"
	"simple_payments/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(username, pasword string) (string, error)
	Logout() error
	Register(username, password string) error
}

type authUseCase struct {
	customerRepo repository.CustomerRepository
}

// Register implements AuthUseCase.
func (a *authUseCase) Register(username string, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed hash password")
		return err
	}

	customer := &entity.Customer{
		Username: username,
		Password: string(hashPassword),
		Balance:  5.0,
	}

	err = a.customerRepo.Register(customer)
	if err != nil {
		return err
	}

	return nil
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(username string, pasword string) (string, error) {
	customer, err := a.customerRepo.FindByUsername(username)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pasword))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	return fmt.Sprint("login succesfull."), nil
}

// Logout implements AuthUseCase.
func (a *authUseCase) Logout() error {
	return nil
}

func NewAuthUseCase(customerRepo repository.CustomerRepository) AuthUseCase {
	return &authUseCase{
		customerRepo: customerRepo,
	}
}
