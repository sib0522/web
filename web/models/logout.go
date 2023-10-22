package models

import (
	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	// セッション取得
	session, _ := Store.Get(c.Request(), "session")

	// ログイン中のアカウントを取得
	val := session.Values["email"]

	if val != nil {
		session.Values["email"] = ""
		Common.LoginEmail = ""
		Common.IsLogin = false
		session.Save(c.Request(), c.Response())
	} else {
		return c.String(404, "ログイン状態ではありません")
	}

	return c.String(200, "ログアウトしました")
}
