package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/pedramcode/goblog/controllers"
)

func UserRouter(group *echo.Group) {
	group.POST("/user/register", controllers.UserRegister)
	group.POST("/user/login", controllers.UserLogin)
}
