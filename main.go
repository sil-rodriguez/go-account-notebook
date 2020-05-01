package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	accountService := createAccountService()
	transactionResource := createTransactionResource()
	transactionResource.transactionService.accountService = accountService
	accountResource := new(AccountResource)
	accountResource.accountService = accountService
	router.HandleFunc("/", accountResource.getBalance).Methods("GET")
	router.HandleFunc("/transactions", transactionResource.makeTransaction).Methods("POST")
	router.HandleFunc("/transactions", transactionResource.findAll).Methods("GET")
	router.HandleFunc("/transactions/{id}", transactionResource.findByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
