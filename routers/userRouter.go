package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/pedramcode/goblog/controllers"
	"github.com/pedramcode/goblog/middlewares"
)

func UserRouter(group *echo.Group) {
	group.POST("/user/register", controllers.UserRegister)
	group.POST("/user/login", controllers.UserLogin)
	group.POST("/user/logout", controllers.UserLogout, middlewares.TokenAuthMiddleware())
	group.POST("/user/me", controllers.UserProfile, middlewares.TokenAuthMiddleware())
}
