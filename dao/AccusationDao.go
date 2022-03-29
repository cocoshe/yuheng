package dao

import (
	"backend/models"
	"github.com/google/uuid"
	"time"
)

func CreatePost(user models.User, post string, picName string) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	accusation := &models.Accusation{
		Id:     uuid,
		UserId: user.Id,
		Time:   time.Now(),
		Post:   post,
	}
	_ = accusation
	//drivers.MysqlDb.Table("accusation").Create(&accusation)
	return nil
}
