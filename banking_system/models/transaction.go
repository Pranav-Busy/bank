package models

import (
   // "time"
)

type Transaction struct {
    TransactionID   uint      `json:"transaction_id" pg:",pk"`  // primary key
    AccountID      uint      `json:"account_id" pg:"on_delete:CASCADE"`      // foreign key referencing Account.ID
    Mode            string    `json:"mode" pg:",notnull"`
    Receiver_accnum  uint    `json:"receiver_account_number"`
	Timestamp       string   `json:"timestamp"`
    Amount          float64   `json:"amount" pg:",notnull"`
	Account        *Account  `json:"account" pg:"rel:has-one"`
}