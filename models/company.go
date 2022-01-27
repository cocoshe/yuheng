package models

type Company struct {
	Id     string `json:"id" gorm:"column:company_id;primary_key" primaryKey:"true"`
	Status int    `json:"status" gorm:"column:status"`
}
