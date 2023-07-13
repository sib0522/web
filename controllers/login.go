package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/login.html")
	return c.Render(http.StatusOK, "login", nil)
}
