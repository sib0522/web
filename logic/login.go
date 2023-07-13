package logic

import (
	"GoEcho/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	account := models.Account{
		Password: password,
		Email:    email,
	}

	id, nickName := account.Read()
	if id == "" || nickName == "" {
		return c.String(http.StatusExpectationFailed, "ログインに失敗しました")
	}

	return c.String(http.StatusOK, "ログインに成功しました")
}
