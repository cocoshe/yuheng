package models

type User struct {
	Id    string `json:"id" primaryKey:"true"`
	Pwd   string `json:"pwd"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

//func (User) TableName() string {
//	return "user"
//}
