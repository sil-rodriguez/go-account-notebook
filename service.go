package main

import (
	"time"

	guuid "github.com/google/uuid"
)

// AccountService is a service used to perform account credit and debit operations
type AccountService struct {
	accountRepository *AccountRepository
}

func createAccountService() *AccountService {
	service := new(AccountService)
	service.accountRepository = createAccountRepository()
	return service
}

func (service AccountService) creditAccount(amount float64) {
	service.accountRepository.creditAccount(amount)
}

func (service AccountService) debitAccount(amount float64) error {
	return service.accountRepository.debitAccount(amount)
}

func (service AccountService) getAccountBalance() (float64, error) {
	return service.accountRepository.getBalance()
}

// TransactionService is a service used to perform transaction operations
type TransactionService struct {
	transactionRepository *TransactionRepository
	accountService        *AccountService
}

func createTransactionService() *TransactionService {
	service := new(TransactionService)
	service.transactionRepository = createTransactionRepository()
	return service
}

func (service TransactionService) makeTransaction(transaction Transaction) (*Transaction, error) {
	if "credit" == transaction.TransactionType {
		service.accountService.creditAccount(transaction.Amount)
	} else {
		err := service.accountService.debitAccount(transaction.Amount)
		if err != nil {
			return nil, err
		}
	}
	transaction.ID = guuid.New().String()
	transaction.Timestamp = time.Now()
	service.transactionRepository.save(transaction)
	return &transaction, nil
}

func (service TransactionService) findByID(id string) (*Transaction, error) {
	return service.transactionRepository.findByID(id)
}

func (service TransactionService) findAll() ([]Transaction, error) {
	return service.transactionRepository.findAll()
}
