package models



type Customer struct {
	ID       uint       `json:"c_id" pg:",pk"` // primary key
	BranchID uint       `json:"branch_id"`     // foreign key referencing Branch.ID
	Name     string     `json:"name" pg:",notnull"`
	PAN      string     `json:"pan" pg:",notnull"`
    DOB      string       `json:"dob" pg:",notnull"`
	Age      uint       `json:"age" pg:",notnull"`
	Phone    string     `json:"phone" pg:",notnull"`
	Address  string     `json:"address" pg:",notnull"`
	Branch   *Branch    `json:"branch" pg:"rel:has-one"`
	Account  []*Account `json:"accounts" pg:"many2many:customer_accounts"`
}
