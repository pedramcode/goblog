package controllers

import (
	"github.com/labstack/echo/v4"
	logic "github.com/pedramcode/goblog/logics"
	"github.com/pedramcode/goblog/models"
	"github.com/pedramcode/goblog/utils"
	"net/http"
)

func UserRegister(ctx echo.Context) error {
	userData := models.User{}

	if err := ctx.Bind(&userData); err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	if userData.Username == "" || userData.Password == "" {
		return utils.RaiseError(&ctx, http.StatusBadRequest, "Pass required parameters")
	}

	user, err := logic.UserCreate(userData.Username, userData.Password, userData.Firstname,
		userData.Lastname, userData.Email, userData.Password)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	token, err := logic.TokenCreate(user.ID)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	res := make(map[string]interface{})
	res["token"] = token.Key
	return utils.StdResponse(&ctx, http.StatusOK, res)
}
