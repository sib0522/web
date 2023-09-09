package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IndexData struct {
	IsLogin       bool
	SideNaviClass string
}

func Index(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/index.html")
	data := IndexData{
		IsLogin:       false,
		SideNaviClass: "sb-sidenav accordion sb-sidenav-light",
	}

	return c.Render(http.StatusOK, "index", &data)
}
