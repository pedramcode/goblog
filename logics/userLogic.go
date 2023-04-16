package logic

import (
	"errors"
	"net/mail"

	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/models"
	utils "github.com/pedramcode/goblog/utils"
)

func UserCreate(username, password, firstname, lastname, email, phone string) (models.User, error) {
	db := core.DB()

	if email != "" {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return models.User{}, err
		}
	}

	obj := models.User{
		Username:  username,
		Password:  utils.HashPassword(password),
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
	}
	res := db.Create(&obj)
	if res.Error != nil {
		return models.User{}, errors.New(res.Error.Error())
	}
	return obj, nil
}

func UserAuthenticate(username string, password string) (models.User, error) {
	db := core.DB()
	obj := models.User{}

	res := db.Where("username LIKE ? AND password LIKE ?", username, utils.HashPassword(password)).First(&obj)
	if res.Error != nil {
		return models.User{}, errors.New(res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return models.User{}, errors.New("User is not authenticated")
	}
	return obj, nil
}

func UserAll() ([]models.User, error) {
	db := core.DB()
	objs := []models.User{}

	res := db.Find(&objs)
	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}
	return objs, nil
}

func UserGetByID(id uint) (models.User, error) {
	db := core.DB()
	obj := models.User{}
	obj.ID = id

	res := db.First(&obj)
	if res.Error != nil {
		return models.User{}, errors.New(res.Error.Error())
	}

	return obj, nil
}

func UserGetByUsername(username string) (models.User, error) {
	db := core.DB()
	obj := models.User{}

	res := db.Where("username LIKE ?", username).First(&obj)
	if res.Error != nil {
		return models.User{}, errors.New(res.Error.Error())
	}

	return obj, nil
}

func UserIsUnique(user *models.User) (bool, error) {
	db := core.DB()
	obj := models.User{}

	res := db.Where("username LIKE ? OR email LIKE ? ", user.Username, user.Email).First(&obj)
	if res.Error != nil {
		return false, errors.New(res.Error.Error())
	}
	return res.RowsAffected == 0, nil
}
