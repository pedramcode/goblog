package utils

import (
	"github.com/labstack/echo/v4"
)

func RaiseError(ctx *echo.Context, status int, message interface{}) error {
	res := make(map[string]interface{})
	res["error"] = message
	return (*ctx).JSON(status, res)
}

func StdResponse(ctx *echo.Context, status int, data interface{}) error {
	res := make(map[string]interface{})
	res["data"] = data
	return (*ctx).JSON(status, res)
}

func StdResponseMsg(ctx *echo.Context, status int, message string, data interface{}) error {
	res := make(map[string]interface{})
	res["data"] = data
	res["msg"] = message
	return (*ctx).JSON(status, res)
}
