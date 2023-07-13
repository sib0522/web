package main

import (
	"GoEcho/router"
	"GoEcho/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// middleware.Loggerブラウザで行った動作がログで表示される
	e.Use(middleware.Logger())
	router.InitRouting(e)
	server.InitServer(e)
}
