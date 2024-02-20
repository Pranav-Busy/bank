package models
type Openacc struct{
	Branch_id   uint    `json:"branch_id" pg:",fk"`        // foreign key referencing Branch.ID
    Balance    float64 `json:"balance" pg:",notnull"`
    Acc_type   string  `json:"account_type" pg:",notnull"`
	Customer []Customer `json:"customer"`
}