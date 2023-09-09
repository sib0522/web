package main

import (
	"GoEcho/app/appRouter"
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
	appRouter.InitAppRouting(e)
	server.InitServer(e)

	//go:generate go run github.com/shamaton/msgpackgen
}
