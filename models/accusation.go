package models

import (
	"github.com/google/uuid"
	"time"
)

type Accusation struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key" example:"a05cde18-488b-4dc4-afa1-df9e0ccf16a1"` //文章uuid
	CompanyId string    `json:"company_id" example:"17280000089799"`
	UserId    string    `json:"userId" example:"123123"`                  //用户名
	Time      time.Time `json:"time" example:"2022-01-28T23:57:30+08:00"` //时间
	Post      string    `json:"post" example:"文字内容"`                      //帖子内容
	Status    int       `json:"status" example:"2"`                       //举报状态
	Pic       string    `json:"pic,omitempty"`                            //图片
	PicType   string    `json:"pic_type,omitempty"`                       //图片类型
}

type Appeal struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key" example:"967d7ae4-1c39-40f8-befd-222d768d70f7"` //文章uuid
	CompanyId string    `json:"company_id" example:"17280000089799"`
	Time      time.Time `json:"time" example:"2022-01-28T23:45:54+08:00"` //时间
	Post      string    `json:"post" example:"文字内容"`                      //帖子内容
	UserId    string    `json:"user_id" example:"上传者用户名"`                 //上传者用户名
	Status    int       `json:"status" example:"2"`                       //举报状态
	Pic       string    `json:"pic,omitempty"`                            //图片地址
	PicType   string    `json:"pic_type,omitempty"`                       //图片类型
}
