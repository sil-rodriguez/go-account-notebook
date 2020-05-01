package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//TransactionResource is a resource for transaction operations
type TransactionResource struct {
	transactionService    *TransactionService
	transactionValidation *TransactionValidation
	idValidation          *IDValidation
}

func createTransactionResource() *TransactionResource {
	resource := new(TransactionResource)
	resource.transactionService = createTransactionService()
	resource.transactionValidation = createTransactionValidation()
	resource.idValidation = new(IDValidation)
	return resource
}

func (resource TransactionResource) makeTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction Transaction
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.Unmarshal(reqBody, &newTransaction)
	err = resource.transactionValidation.validate(newTransaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	transaction, err := resource.transactionService.makeTransaction(newTransaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func (resource TransactionResource) findByID(w http.ResponseWriter, r *http.Request) {
	transactionID := mux.Vars(r)["id"]
	err := resource.idValidation.validate(transactionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	transaction, err := resource.transactionService.findByID(transactionID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

func (resource TransactionResource) findAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := resource.transactionService.findAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

//AccountResource is a resource for account balance operations
type AccountResource struct {
	accountService *AccountService
}

func (accountResource AccountResource) getBalance(w http.ResponseWriter, r *http.Request) {
	balance, err := accountResource.accountService.getAccountBalance()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}
