package models

type User struct {
	Id     string `json:"id" gorm:"primary_key"`            // 账号
	Pwd    string `json:"pwd"`                              // 密码
	Name   string `json:"name"`                             // 昵称/姓名
	Level  int    `json:"level"`                            // 权限
	Reward int    `json:"reward"`                           //积分
	Atten  string `json:"atten" gorm:"type:varchar(10240)"` //关注
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
