package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func InitServer(e *echo.Echo) {
	fmt.Println("____________[Init]Server____________")
	e.Logger.Fatal(e.Start(":1323"))
}
