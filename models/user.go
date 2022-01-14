package models

type User struct {
	Id    string `json:"id" primaryKey:"true"` // 账号
	Pwd   string `json:"pwd"`                  // 密码
	Name  string `json:"name"`                 // 昵称/姓名
	Level int    `json:"level"`                // 权限
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
