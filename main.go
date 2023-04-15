package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/routers"
)

func main() {
	core.InitDb()
	core.Migrate()

	server := echo.New()
	apiGroup := server.Group("/api")

	routers.UserRouter(apiGroup)
	routers.PostRouter(apiGroup)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Logger.Fatal(server.Start(":8080"))
}
