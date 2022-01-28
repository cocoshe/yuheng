package models

import (
	"github.com/google/uuid"
	"time"
)

type Accusation struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"` //文章uuid
	CompanyId string    `json:"company_id"`
	UserId    string    `json:"userId"` //用户名
	Time      time.Time `json:"time"`   //时间
	Post      string    `json:"post"`   //帖子内容
	Pic       string    `json:"pic"`    //帖子图片路径
	Status    int       `json:"status"` //举报状态
}

type Appeal struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"` //文章uuid
	CompanyId string    `json:"company_id"`
	Time      time.Time `json:"time"`   //时间
	Post      string    `json:"post"`   //帖子内容
	Pic       string    `json:"pic"`    //帖子图片路径
	Status    int       `json:"status"` //举报状态
}
