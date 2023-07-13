package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/index.html")
	return c.Render(http.StatusOK, "index", nil)
}
