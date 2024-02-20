package models


type Bank struct {
    ID       uint               `json:"id" pg:",pk"`
    Name     string             `json:"name" pg:",notnull"`
    Branches []*Branch   `json:"branches" pg:"rel:has-many"`
}
