package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TokenAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, ctx echo.Context) (bool, error) {
		return true, nil
	})
}
