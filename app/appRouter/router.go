package appRouter

import (
	"GoEcho/app/controllers"
	"github.com/labstack/echo/v4"
)

func InitAppRouting(e *echo.Echo) {
	e.POST("/user/:apiName", controllers.UserController)
}
