package models

type User struct {
	Id    string `json:"id" from:"id" primaryKey:"true"`
	Pwd   string `json:"pwd" from:"pwd" primaryKey:"required"`
	Name  string `json:"name" from:"name" primaryKey:"required"`
	Level int    `json:"level" from:"level" primaryKey:"required"`
}

//func (User) TableName() string {
//	return "user"
//}
