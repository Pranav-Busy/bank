package models

type Account struct {
    ID         uint    `json:"account_id" pg:",pk"`       // primary key
    Branch_id   uint    `json:"branch_id" pg:"on_delete:CASCADE"`        // foreign key referencing Branch.ID
	AccNumber  uint    `json:"accnumber" pg:",notnull,unique,default:nextval('account_id_seq')"`
    Balance    float64 `json:"balance" pg:",notnull"`
    Acc_type   string  `json:"account_type" pg:",notnull"`
	Branch    *Branch  `json:"branch" pg:"rel:has-one"` 
	Transaction []*Transaction `json:"transaction" pg:"rel:has-many"` 
	Customer []*Customer `json:"customer" pg:"many2many:customer_accounts"`
    
}