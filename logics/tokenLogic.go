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

func TokenGetByUserIDCreate(userID uint) (models.Token, error) {
	db := core.DB()
	obj := models.Token{}
	obj.UserID = userID

	res := db.First(&obj)
	if res.RowsAffected == 0 {
		return TokenCreate(userID)
	}
	if res.Error != nil {
		return models.Token{}, errors.New(res.Error.Error())
	}

	return obj, nil
}

func TokenGetByKey(key string) (models.Token, error) {
	db := core.DB()
	obj := models.Token{}

	res := db.Where("Key LIKE ?", key).First(&obj)
	if res.Error != nil {
		return models.Token{}, errors.New(res.Error.Error())
	}

	return obj, nil
}

func TokenDeleteByUserID(userID uint) error {
	db := core.DB()
	obj := models.Token{}

	res := db.Unscoped().Where("user_id = ?", userID).Delete(&obj)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}
	return nil
}
