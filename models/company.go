package models

type Company struct {
	Id     string `json:"id" gorm:"column:company_id;primary_key" primaryKey:"true" example:"172800010800110000"`
	Status int    `json:"status,omitempty" gorm:"column:status" example:"2"`
}

type CompanyDto struct {
	Id     string `json:"id" gorm:"column:company_id;primary_key" primaryKey:"true" example:"172800010800110000"`
	Status int    `json:"status" gorm:"column:status" example:"2"`
}
