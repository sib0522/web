package controllers

import (
	"GoEcho/forms"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/delete.html")
	return c.Render(http.StatusOK, "delete", nil)
}
