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
