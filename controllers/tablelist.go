package controllers

import (
	"GoEcho/forms"
	"GoEcho/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TableList(c echo.Context) error {
	if c.Path() == "/gamedbtablelist" {
		forms.SetRenderer(c.Echo(), "views/gamedbtablelist.html")
		return c.Render(http.StatusOK, "gamedbtablelist", nil)
	} else {
		forms.SetRenderer(c.Echo(), "views/tablelist.html")
		return c.Render(http.StatusOK, "tablelist", nil)
	}
}

func TableDetail(c echo.Context) error {
	name := c.Param("name")
	forms.SetRenderer(c.Echo(), "views/tabledetail.html")
	return c.Render(http.StatusOK, "tabledetail", logic.CreateTable(name))
}
