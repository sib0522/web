package controllers

import (
	"GoEcho/forms"
	"GoEcho/logic"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TableList(c echo.Context) error {
	_, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}

	if c.Path() == "/table/game" {
		forms.SetRenderer(c.Echo(), "views/table_list_game.html")
		return c.Render(http.StatusOK, "table_list_game", nil)
	} else {
		forms.SetRenderer(c.Echo(), "views/table_list_admin.html")
		return c.Render(http.StatusOK, "table_list_admin", nil)
	}
}

func TableDetail(c echo.Context) error {
	name := c.Param("name")
	forms.SetRenderer(c.Echo(), "views/table_detail.html")
	return c.Render(http.StatusOK, "table_detail", logic.CreateTable(name))
}
