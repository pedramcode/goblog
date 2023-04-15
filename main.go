package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pedramcode/goblog/core"
	"github.com/pedramcode/goblog/routers"
)

func main() {
	core.Init_DB()
	core.Migrate()

	server := echo.New()
	apiGroup := server.Group("/api")
	routers.UserRouter(apiGroup)
	routers.PostRouter(apiGroup)
	server.Logger.Fatal(server.Start(":8080"))
}
