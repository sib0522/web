package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	ok, _ := IsLogin(c)
	if !ok {
		return c.String(http.StatusUnauthorized, "ログイン状態ではありません")
	}

	err := SessionLogout(c)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "ログアウトしました")
}
