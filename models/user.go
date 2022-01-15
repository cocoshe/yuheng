package models

type User struct {
	Id     string `json:"id" gorm:"id" primaryKey:"true"` // 账号
	Pwd    string `json:"pwd" gorm:"pwd"`                 // 密码
	Name   string `json:"name" gorm:"name"`               // 昵称/姓名
	Level  int    `json:"level" gorm:"level"`             // 权限
	Reward int    `json:"reward" gorm:"reward"`           //积分
	status int    `json:"status" gorm:"status"`
}

/*
example:
{
    "id": "20041329",
    "pwd": "123",
    "name": "cocoshe",
    "level": 100
}
*/

//func (User) TableName() string {
//	return "user"
//}
