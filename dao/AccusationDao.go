package dao

import (
	"backend/models"
	"fmt"
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
		Pic:    fmt.Sprintf("static/img/%s", uuid),
	}
	_ = accusation
	//drivers.MysqlDb.Table("accusation").Create(&accusation)
	return nil
}
