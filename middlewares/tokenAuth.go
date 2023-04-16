package middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logic "github.com/pedramcode/goblog/logics"
)

func TokenAuthMiddleware() echo.MiddlewareFunc {
	// TODO Not optimized query
	return middleware.KeyAuth(func(key string, ctx echo.Context) (bool, error) {
		token, err := logic.TokenGetByKey(key)
		if err != nil {
			return false, errors.New(err.Error())
		}
		user, err := logic.UserGetByID(token.UserID)
		if err != nil {
			return false, errors.New(err.Error())
		}
		ctx.Set("user", user)
		return true, nil
	})
}
