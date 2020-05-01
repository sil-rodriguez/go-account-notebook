package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestReturnBalanceWhenValidTransactions(t *testing.T) {
	accountRepository := new(AccountRepository)
	accountRepository.creditAccount(5)
	accountRepository.debitAccount(3)
	balance, _ := accountRepository.getBalance()
	assert.EqualValues(t, 2, balance)
}

func TestReturnErrorWhenInsufficientBalance(t *testing.T) {
	accountRepository := new(AccountRepository)
	accountRepository.creditAccount(1)
	err := accountRepository.debitAccount(3)
	assert.EqualValues(t, "Account balance is insufficient", err.Error())
}

func TestReturnTransactions(t *testing.T) {
	transactionRepository := createTransactionRepository()
	transaction1 := new(Transaction)
	transaction1.ID = "id1"
	transaction2 := new(Transaction)
	transaction1.ID = "id2"
	transactionRepository.save(*transaction1)
	transactionRepository.save(*transaction2)
	transactions, _ := transactionRepository.findAll()
	assert.EqualValues(t, 2, len(transactions))
}

func TestReturnTransactionWhenValidTransactionID(t *testing.T) {
	transactionRepository := createTransactionRepository()
	transaction1 := new(Transaction)
	transaction1.ID = "id1"
	transactionRepository.save(*transaction1)
	transaction, _ := transactionRepository.findByID("id1")
	assert.NotNil(t, transaction)
}

func TestReturnErrorWhenTransactionIDNotFound(t *testing.T) {
	transactionRepository := createTransactionRepository()
	_, err := transactionRepository.findByID("id1")
	assert.EqualValues(t, "No transaction was found for id: id1", err.Error())
}
