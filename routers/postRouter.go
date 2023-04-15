package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/pedramcode/goblog/controllers"
	"github.com/pedramcode/goblog/middlewares"
)

func PostRouter(group *echo.Group) {
	group.POST("/post", controllers.PostCreate, middlewares.TokenAuthMiddleware())
	group.GET("/post/me", controllers.PostList, middlewares.TokenAuthMiddleware())
	group.DELETE("/post/:id", controllers.PostDelete, middlewares.TokenAuthMiddleware())
	group.GET("/post", controllers.PostPubList)
}
