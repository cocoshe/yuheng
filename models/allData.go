package models

type AllData struct {
	CompanyId     string  `json:"company_id" gorm:"column:company_id;primary_key" primaryKey:"true" example:"172800010800110000"` // 公司id
	PortId        string  `json:"port_id" gorm:"column:port_id" example:"64000000000600100000 "`                                  // 排污口id
	PolutionId    string  `json:"polution_id" gorm:"column:polution_id" example:"w00000 "`                                        // 污染物id
	Date          string  `json:"date" gorm:"column:date" example:"2021-11-30 00:00:00"`                                          // 日期
	Concentration float64 `json:"concentration" gorm:"column:concentration" example:"0.0"`                                        // 浓度
	Amount        float64 `json:"amount" gorm:"column:amount" example:"0.0"`                                                      // 排放量
}
