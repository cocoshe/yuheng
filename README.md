# yuheng


+ Tables:

```go
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


type User struct {
Id     string `json:"id" gorm:"id" primaryKey:"true"` // 账号
Pwd    string `json:"pwd" gorm:"pwd"`                 // 密码
Name   string `json:"name" gorm:"name"`               // 昵称/姓名
Level  int    `json:"level" gorm:"level"`             // 权限
Reward int    `json:"reward" gorm:"reward"`           //积分
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

```