package main

import "sync"

// TransactionRepository is a repository used to handle transaction save and retrieve operations
type TransactionRepository struct {
	transactions map[string]Transaction
}

func createTransactionRepository() *TransactionRepository {
	repository := new(TransactionRepository)
	repository.transactions = make(map[string]Transaction)
	return repository
}

func (repo *TransactionRepository) save(transaction Transaction) {
	repo.transactions[transaction.ID] = transaction
}

func (repo *TransactionRepository) findByID(id string) (*Transaction, error) {
	transaction, ok := repo.transactions[id]
	if ok {
		return &transaction, nil
	}
	notFoundError := new(TransactionNotFoundError)
	notFoundError.message = "No transaction was found for id: " + id
	return nil, notFoundError
}

func (repo *TransactionRepository) findAll() ([]Transaction, error) {
	values := make([]Transaction, 0, len(repo.transactions))
	for _, v := range repo.transactions {
		values = append(values, v)
	}
	return values, nil
}

// AccountRepository is a repository used to perform account credit and debit operations
type AccountRepository struct {
	balance       float64
	readWriteLock *sync.RWMutex
}

func createAccountRepository() *AccountRepository {
	repository := new(AccountRepository)
	repository.readWriteLock = new(sync.RWMutex)
	return repository
}

func (accountRepo *AccountRepository) creditAccount(credit float64) error {
	accountRepo.readWriteLock.Lock()
	accountRepo.balance = accountRepo.balance + credit
	accountRepo.readWriteLock.Unlock()
	return nil
}

func (accountRepo *AccountRepository) debitAccount(debit float64) error {
	accountRepo.readWriteLock.Lock()
	if accountRepo.balance >= debit {
		accountRepo.balance = accountRepo.balance - debit
		accountRepo.readWriteLock.Unlock()
		return nil
	}
	accountRepo.readWriteLock.Unlock()
	insufficentBalanceError := new(InsufficientBalanceError)
	insufficentBalanceError.message = "Account balance is insufficient"
	return insufficentBalanceError
}

func (accountRepo *AccountRepository) getBalance() (float64, error) {
	accountRepo.readWriteLock.RLock()
	balance := accountRepo.balance
	accountRepo.readWriteLock.RUnlock()
	return balance, nil
}
