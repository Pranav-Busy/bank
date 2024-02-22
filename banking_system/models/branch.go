package models

import (
	"github.com/google/uuid"
)

type Branch struct {
    ID      uint   `json:"id" pg:",pk"`
    Name    string `json:"name" pg:",notnull"`
    BankID  uint   `json:"bank_id" pg:"on_delete:CASCADE"`
    Bank    *Bank  `json:"bank" pg:"rel:has-one"` 
	Ifsc_code uuid.UUID `json:"ifsc_code" pg:"type:uuid,default:uuid_generate_v4()"`
    Customer []*Customer   `json:"customers" pg:"rel:has-many"`
    Account  []*Account  `json:"accounts" pg:"rel:has-many"`
}
