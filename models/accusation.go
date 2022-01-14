package models

import (
	"github.com/google/uuid"
	"time"
)

type Accusation struct {
	Id     uuid.UUID `json:"id" gorm:"id"`         //文章uuid
	UserId string    `json:"userId" gorm:"userId"` //用户名
	Time   time.Time `json:"time" gorm:"time"`     //时间
	Post   string    `json:"post" gorm:"post"`     //帖子内容
	Pic    string    `json:"pic" gorm:"pic"`       //帖子图片路径
	Status int       `json:"status" gorm:"status"` //举报状态
}

type Appeal struct {
	Accusation
}
