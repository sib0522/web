package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Password(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/password.html")
	return c.Render(http.StatusOK, "password", nil)
}
