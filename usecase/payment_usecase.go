package usecase

import (
	"errors"
	"fmt"
	"simple_payments/entity"
	"simple_payments/repository"
)

var ErrCustomerNotFound = errors.New("customer not found")

type PaymentUseCase interface {
	MakePayment(sender, receiver string, amount float64) error
	GetTransactionHistory() ([]entity.Transaction, error)
}

type paymentUseCase struct {
	customerRepo repository.CustomerRepository
	merchantRepo repository.MerchantRepository
	historyRepo  repository.TransactionHistoryRepository
}

func (a *paymentUseCase) MakePayment(sender string, receiver string, amount float64) error {
	senderCustomer, err := a.customerRepo.FindByUsername(sender)
	fmt.Println(senderCustomer, "ini sender")
	if err != nil {
		fmt.Println(err)
		if err == ErrCustomerNotFound {
			return fmt.Errorf("sender not found")
		} else {
			// handle other types of errors
			return fmt.Errorf("unknown error occurred: %v", err)
		}
	}

	receiverMerchant, err := a.merchantRepo.FindByName(receiver)
	fmt.Println(receiverMerchant, "ini receiver")
	if err != nil {
		fmt.Println("Error fetching receiver:", err)
		return fmt.Errorf("receiver not found")
	}

	if senderCustomer.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	senderCustomer.Balance -= amount
	receiverMerchant.Balance += amount

	transaction := entity.Transaction{From: sender, To: receiver, Amount: amount}
	err = a.historyRepo.SaveTransaction(&transaction)
	if err != nil {
		return fmt.Errorf("failed to save transaction to history")
	}

	err = a.customerRepo.Save(senderCustomer)
	if err != nil {
		return fmt.Errorf("failed to update sender balance")
	}

	err = a.merchantRepo.SaveMerchant(receiverMerchant)
	if err != nil {
		return fmt.Errorf("failed to update receiver balance")
	}

	return nil
}

func (a *paymentUseCase) GetTransactionHistory() ([]entity.Transaction, error) {
	return a.historyRepo.GetAll()
}

func NewPaymentUsecase(customerRepo repository.CustomerRepository, merchantRepo repository.MerchantRepository, historyRepo repository.TransactionHistoryRepository) PaymentUseCase {
	return &paymentUseCase{
		customerRepo: customerRepo,
		merchantRepo: merchantRepo,
		historyRepo:  historyRepo,
	}
}
