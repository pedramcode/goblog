package main

import (
	"github.com/pedramcode/goblog/core"
)

func main() {
	core.Init_DB()
	core.Migrate()
}
