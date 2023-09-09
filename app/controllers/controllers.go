package controllers

import (
	"GoEcho/app/application"
	"github.com/labstack/echo/v4"
)

func UserController(c echo.Context) error {
	api := c.Request().RequestURI
	switch api {
	case "/user/create":
		s := application.NewUserCreateService()
	case "/user/update":
	case "/user/load":
	}

	return nil
}
