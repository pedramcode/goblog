package logic

import (
	"errors"
	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/models"
)

func PostCreate(userID uint, title, content string, comments []models.Comment) (models.Post, error) {
	db := core.DB()
	obj := models.Post{
		UserID:   userID,
		Title:    title,
		Content:  content,
		Comments: comments,
	}

	res := db.Create(&obj)
	if res.Error != nil {
		return models.Post{}, errors.New(res.Error.Error())
	}

	return obj, nil
}

func PostGetByUserID(userID uint) ([]models.Post, error) {
	db := core.DB()
	var objs []models.Post

	res := db.Where("user_id = ?", userID).Find(&objs)
	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}
	return objs, nil
}

func PostGetAll() ([]models.Post, error) {
	db := core.DB()
	var objs []models.Post
	res := db.Find(&objs)
	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}
	return objs, nil
}

func PostDeleteByID(userID uint, postID uint) error {
	db := core.DB()
	var objs []models.Post

	res := db.Where("user_id = ? AND id = ?", userID, postID).Delete(&objs)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}
	return nil
}
