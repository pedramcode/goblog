package logic

import (
	"errors"
	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/models"
)

func CommentCreate(postID uint, userID uint, content string) (models.Comment, error) {
	db := core.DB()
	obj := models.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	res := db.Create(&obj)
	if res.Error != nil {
		return models.Comment{}, errors.New(res.Error.Error())
	}
	return obj, nil
}
