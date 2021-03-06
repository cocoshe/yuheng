package dao

import (
	"backend/drivers"
	"backend/models"
	"errors"
)

func CreateUser(user *models.User) error {
	if user.Pwd == "" || user.Id == "" || user.Name == "" {
		return errors.New("Create new user failed!")
	}
	drivers.MysqlDb.Table("user").Create(user)
	return nil
}

func DeleteUser(user *models.User) error {
	if user.Id == "" {
		return errors.New("Id is empty, delete failed!")
	}
	drivers.MysqlDb.Table("user").Delete(user)
	return nil
}

func DeleteUserById(id string) error {
	if id == "" {
		return errors.New("Id is empty, delete failed!")
	}
	user := &models.User{}
	FindUserById(id, user)
	drivers.MysqlDb.Table("user").Delete(user)
	return nil
}

func FindUserById(id string, user *models.User) error {
	if id == "" {
		return errors.New("id is empty, find failed!")
	}
	if user != nil {
		drivers.MysqlDb.Table("user").Where("id = ?", id).First(user)
	}
	return nil
}

func UpdateUser(user *models.User) error {
	if user.Id == "" {
		return errors.New("userId is empty, update failed!")
	}
	drivers.MysqlDb.Table("user").Save(user)
	return nil
}
