package main

import "time"

//Transaction is a model to represent transactions
type Transaction struct {
	ID              string    `json:"id"`
	TransactionType string    `json:"type"`
	Amount          float64   `json:"amount"`
	Timestamp       time.Time `json:"timestamp"`
}
