package controllers

import (
	"github.com/labstack/echo/v4"
	logic "github.com/pedramcode/goblog/logics"
	"github.com/pedramcode/goblog/models"
	"github.com/pedramcode/goblog/utils"
	"net/http"
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
