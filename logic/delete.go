package logic

import (
	"GoEcho/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	ses, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	if b, err := ses.Values["in"]; b == false || err {
		return c.String(http.StatusUnauthorized, "401")
	}
	acc := &models.Account{
		Id: ses.Values["id"].(uint),
	}
	if b := acc.Delete(); !b {
		return c.String(http.StatusInternalServerError, "アカウントの削除に失敗しました")
	}

	return c.String(http.StatusOK, "アカウントを削除しました")
}
