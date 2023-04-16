package controllers

import (
	"github.com/labstack/echo/v4"
	logic "github.com/pedramcode/goblog/logics"
	"github.com/pedramcode/goblog/models"
	"github.com/pedramcode/goblog/pipelines"
	"github.com/pedramcode/goblog/serializers"
	"github.com/pedramcode/goblog/utils"
	"net/http"
	"strconv"
)

func PostCreate(ctx echo.Context) error {
	user, _ := ctx.Get("user").(models.User)

	postData := models.Post{}
	if err := ctx.Bind(&postData); err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	if postData.Title == "" || postData.Content == "" {
		return utils.RaiseError(&ctx, http.StatusBadRequest, "Body or title cannot be empty")
	}

	post, err := logic.PostCreate(user.ID, postData.Title, postData.Content, nil)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	return utils.StdResponse(&ctx, http.StatusOK, post)
}

func PostList(ctx echo.Context) error {
	user, _ := ctx.Get("user").(models.User)
	posts, err := logic.PostGetByUserID(user.ID)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	return utils.StdResponse(&ctx, http.StatusOK, posts)
}

func PostDelete(ctx echo.Context) error {
	user, _ := ctx.Get("user").(models.User)
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	err = logic.PostDeleteByID(user.ID, uint(idInt))
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	return utils.StdResponse(&ctx, http.StatusOK, "Post deleted")
}

func PostPubList(ctx echo.Context) error {
	posts, err := logic.PostGetAll()
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	return utils.StdResponse(&ctx, http.StatusOK, posts)
}

func PostNewComment(ctx echo.Context) error {
	user, _ := ctx.Get("user").(models.User)

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	commentData := serializers.NewCommentSerializer{}
	if err := ctx.Bind(&commentData); err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	if commentData.Content == "" {
		return utils.RaiseError(&ctx, http.StatusBadRequest, "Content cannot be empty")
	}

	wordChannel := pipelines.WordChannelize(commentData.Content)
	junkChannel := pipelines.JunkChannelize(wordChannel)
	for jRes := range junkChannel {
		if jRes {
			return utils.RaiseError(&ctx, http.StatusBadRequest, "Comment contains banned words")
		}
	}

	comment, err := logic.CommentCreate(uint(idInt), user.ID, commentData.Content)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	return utils.StdResponse(&ctx, http.StatusOK, comment)
}
