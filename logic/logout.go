package logic

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	ses, _ := session.Get("session", c)
	ses.Values["in"] = false

	if err := ses.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.String(http.StatusInternalServerError, "ログアウトしました")
}
