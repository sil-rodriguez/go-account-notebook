package main

//ValidationError is an error to express that a request format is not valid
type ValidationError struct {
	message string
}

func (validationError *ValidationError) Error() string {
	return validationError.message
}

//InsufficientBalanceError is an error to express that an account does not have enough balance to perform an operation
type InsufficientBalanceError struct {
	message string
}

func (insufficientBalanceError *InsufficientBalanceError) Error() string {
	return insufficientBalanceError.message
}

//TransactionNotFoundError is an error to express that a transaction does not exist
type TransactionNotFoundError struct {
	message string
}

func (transactionNotFoundError *TransactionNotFoundError) Error() string {
	return transactionNotFoundError.message
}
