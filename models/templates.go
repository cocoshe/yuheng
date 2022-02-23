package models

import "github.com/google/uuid"

type DisplayRequest struct {
	Idx1 string `json:"idx1" example:"1"`
	Idx2 string `json:"idx2" example:"3"`
}

type UserPostDto struct {
	Id   string `json:"id" example:"17280000089799"`
	Post string `json:"post,omitempty" example:"文字内容"`
	Pic  string `json:"pic" example:"图片内容"`
}

type CompanyAppealDto struct {
	Id   string `json:"id" example:"17280000089799"`
	Post string `json:"post" example:"文字申诉内容"`
	Pic  string `json:"pic,omitempty" example:"图片内容"`
}

type AccusStatusDto struct {
	Id     uuid.UUID `json:"id" example:"9cade01e-4c66-4c57-910a-2efe5f7e3dff"`
	Status int       `json:"status" example:"1"`
}

type AplStatusDto struct {
	Id     uuid.UUID `json:"id" example:"967d7ae4-1c39-40f8-befd-222d768d70f7"`
	Status int       `json:"status" example:"1"`
}

type SuccessResponse struct {
	Code string `json:"code" example:"200"`
	Msg  string `json:"msg" example:"success"`
}

type AccusationList struct {
	Code string       `json:"code" example:"200"`
	Msg  string       `json:"msg" example:"success"`
	Data []Accusation `json:"data"`
}

type AppealList struct {
	Code string   `json:"code" example:"200"`
	Msg  string   `json:"msg" example:"success"`
	Data []Appeal `json:"data"`
}

type FailureResponse struct {
	Code string `json:"code" example:"401"`
	Msg  string `json:"msg" example:"token invalid"`
}

type LoginRequest struct {
	Id  string `json:"id" gorm:"primary_key" example:"321321"` // 账号
	Pwd string `json:"pwd" example:"321"`                      // 密码
}

type LoginResponse struct {
	Code  string `json:"code" example:"200" example:"200"`
	Msg   string `json:"msg" example:"success" example:"success"`
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE3MjgwMDAwMDg5OSIsImlhdCI6MTU4NjQwMjQ5NywiZXhwIjoxNTg2NDAyNzk3fQ."`
}

type DisplayResponse struct {
	Code string    `json:"code" example:"200"`
	Msg  string    `json:"msg" example:"success"`
	Data []Company `json:"data"`
}

type JSONid struct {
	Id string `json:"id" example:"17280000089799"`
}

type CpnListResponse struct {
	Code string    `json:"code" example:"200"`
	Msg  string    `json:"msg" example:"success"`
	Data []Company `json:"data"`
}

// todo: 完成另一个时间段的采取, 跑两个维度的数据
type RunModelRequest struct {
	PortId     string `json:"port_id" example:"64000000000600100000"` // 污染口ID
	PolutionId string `json:"polution_id" example:"w00000"`           // 污染物ID
	CompanyId  string `json:"company_id" example:"17280000089583"`    // 公司ID
	DataS1     string `json:"date_s_1" example:"1/1/2020"`            // 日/月/年,当前开始时间
	DataE1     string `json:"date_e_1" example:"1/2/2020"`            // 日/月/年,当前结束时间
	DataS2     string `json:"date_s_2" example:"1/3/2020"`            // 日/月/年,对照数据开始时间
	DataE2     string `json:"date_e_2" example:"1/4/2020"`            // 日/月/年,对照数据结束时间
	Threshold  string `json:"threshold" example:"100"`                // 阈值
}

type RunModelResponse struct {
	Code string           `json:"code" example:"200"`
	Msg  string           `json:"msg" example:"success"`
	Data RunModelRespData `json:"data"`
}

// TODO: 跑完模型后, 返回数据
type RunModelRespData struct {
	// 总览
	DangerNum       int `json:"danger_num" example:"1"`         // 疑似违规天数
	CompareOtherNum int `json:"compare_other_num" example:"+1"` // 对照其他公司天数

	// 具体污染物
	AverageConcentration   float64 `json:"average_concentration" example:"0.2"`      // 平均浓度
	CompareThreshold       float64 `json:"compare_threshold" example:"-0.2"`         // 较阈值
	CompareOtherChange     float64 `json:"compare_other_change" example:"0.2"`       // 相对对照数据的变化
	CompareOtherChangeRate float64 `json:"compare_other_change_rate" example:"+0.2"` // 相对对照数据的变化率
	AverageAmount          float64 `json:"average_amount" example:"100"`             // 平均量
	DangerPolutionNum      int     `json:"danger_polution_num" example:"1"`          // 对该污染物的疑似违规天数

	// 可视化
	ComparedDataAverageConcentration float64 `json:"compared_data_average_concentration" example:"2.5"` // 如 上周平均浓度
	ComparedDataAverageAmount        float64 `json:"compared_data_average_amount" example:"100"`        // 如 上周平均量
	// 来源排污口(对应一个排污口)
	PortId            string   `json:"port_id" example:"64000000000600100000"`    // 污染口ID
	PortAmount        float64  `json:"port_amount" example:"1000"`                // 该排污口排污量
	PortConcentration float64  `json:"port_concentration" example:"3.5"`          // 平均排污浓度
	PortRelativeCpn   []string `json:"port_relative_cpn" example:"123, 456, 789"` // 其他共用企业id
}
