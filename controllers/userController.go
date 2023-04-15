package controllers

import (
	"github.com/labstack/echo/v4"
	logic "github.com/pedramcode/goblog/logics"
	"github.com/pedramcode/goblog/models"
	"github.com/pedramcode/goblog/serializers"
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

func UserLogin(ctx echo.Context) error {
	loginData := serializers.LoginSerializer{}

	if err := ctx.Bind(&loginData); err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	user, err := logic.UserAuthenticate(loginData.Username, loginData.Password)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	token, err := logic.TokenGetByUserIDCreate(user.ID)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}

	res := make(map[string]interface{})
	res["token"] = token.Key
	return utils.StdResponse(&ctx, http.StatusOK, res)
}

func UserLogout(ctx echo.Context) error {
	user, _ := ctx.Get("user").(models.User)
	err := logic.TokenDeleteByUserID(user.ID)
	if err != nil {
		return utils.RaiseError(&ctx, http.StatusBadRequest, err.Error())
	}
	return nil
}
