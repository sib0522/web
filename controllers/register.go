package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/register.html")
	return c.Render(http.StatusOK, "register", nil)
}
