package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/index.html")
	var islogin bool

	ses, _ := session.Get("session", c)
	if ses == nil {
		islogin = false
	} else {
		islogin = ses.Values["in"].(bool)
	}

	return c.Render(http.StatusOK, "index", map[string]bool{
		"isLogin": islogin,
	})
}
