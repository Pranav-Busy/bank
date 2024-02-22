package models

type CustomerAccount struct {
	ID         uint
	AccountID  uint
	Account    *Account `json:"accounts" pg:"rel:has-one"`
	CustomerID uint
	Customer   *Customer `json:"customers" pg:"rel:has-one"`
}
