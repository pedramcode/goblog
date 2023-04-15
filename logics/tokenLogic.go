package logic

import (
	"errors"
	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/models"
	"github.com/pedramcode/goblog/utils"
)

func TokenCreate(userID uint) (models.Token, error) {
	db := core.DB()
	token := models.Token{
		UserID: userID,
		Key:    utils.GenerateToken(16),
	}
	res := db.Create(&token)
	if res.Error != nil {
		return models.Token{}, errors.New(res.Error.Error())
	}
	return token, nil
}

func TokenGetByUserID(userID uint) (models.Token, error) {
	db := core.DB()
	obj := models.Token{}
	obj.UserID = userID

	res := db.First(&obj)
	if res.Error != nil {
		return models.Token{}, errors.New(res.Error.Error())
	}

	return obj, nil
}
