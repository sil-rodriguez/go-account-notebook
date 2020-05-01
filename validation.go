package main

import (
	"strings"

	guuid "github.com/google/uuid"
)

//TransactionValidation is a validation for the transaction request format
type TransactionValidation struct {
	transactionTypes []string
}

func createTransactionValidation() *TransactionValidation {
	validation := new(TransactionValidation)
	validation.transactionTypes = []string{"credit", "debit"}
	return validation
}

func (validation TransactionValidation) validate(transaction Transaction) error {
	errorMessages := make([]string, 0)
	if transaction.Amount <= 0 {
		errorMessages = append(errorMessages, "Transaction amount must be greater than 0")
	}
	if !isFound(validation.transactionTypes, transaction.TransactionType) {
		errorMessages = append(errorMessages, "Transaction type must be 'credit' or 'debit'")
	}
	if len(errorMessages) > 0 {
		validationError := new(ValidationError)
		validationError.message = strings.Join(errorMessages, ",")
		return validationError
	}
	return nil
}

func isFound(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

//IDValidation is a validation for the transaction id format
type IDValidation struct {
}

func (validation IDValidation) validate(id string) error {
	_, error := guuid.Parse(id)
	if error != nil {
		validationError := new(ValidationError)
		validationError.message = "Transaction id must be a valid UUID"
		return validationError
	}
	return nil
}
