package main

import (
	"testing"

	guuid "github.com/google/uuid"
	assert "github.com/stretchr/testify/assert"
)

func TestReturnNilWhenValidID(t *testing.T) {
	validation := new(IDValidation)
	id, err := guuid.NewUUID()
	err = validation.validate(id.String())
	assert.Nil(t, err)
}

func TestReturnErrorWhenInvalidID(t *testing.T) {
	validation := new(IDValidation)
	err := validation.validate("a invalid id")
	assert.EqualValues(t, "Transaction id must be a valid UUID", err.Error())
}

func TestReturnNilWhenValidTransaction(t *testing.T) {
	validation := createTransactionValidation()
	transaction := new(Transaction)
	transaction.Amount = 1
	transaction.TransactionType = "credit"
	err := validation.validate(*transaction)
	assert.Nil(t, err)
}

func TestReturnErrorWhenValidTransactionAmout(t *testing.T) {
	validation := createTransactionValidation()
	transaction := new(Transaction)
	transaction.TransactionType = "credit"
	err := validation.validate(*transaction)
	assert.EqualValues(t, "Transaction amount must be greater than 0", err.Error())
}

func TestReturnErrorWhenValidTransactionType(t *testing.T) {
	validation := createTransactionValidation()
	transaction := new(Transaction)
	transaction.Amount = 1
	err := validation.validate(*transaction)
	assert.EqualValues(t, "Transaction type must be 'credit' or 'debit'", err.Error())
}
