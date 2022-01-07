package models

type User struct {
	Id    int    `json:"id" from:"id" primaryKey:"true"`
	Pwd   string `json:"pwd" from:"pwd" primaryKey:"required"`
	Name  string `json:"name" from:"name" primaryKey:"required"`
	Level int    `json:"age" from:"age" primaryKey:"required"`
}
