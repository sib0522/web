package logic

import (
	"GoEcho/models"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// Viewからemailとpasswordを取得する
	email := c.FormValue("email")
	password := c.FormValue("password")

	account := &models.Account{
		Password: password,
		Email:    email,
	}

	// 読み取りできるアカウントのデータがないと一致するデータが存在しないことにする
	b := account.Read()
	if b == false {
		return c.String(http.StatusOK, "ログインに失敗しました")
	}

	c.Echo().Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	ses, _ := session.Get("session", c)
	ses.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	ses.Values["id"] = account.Id
	ses.Values["nickName"] = account.Nickname
	ses.Values["in"] = true

	if err := ses.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "ログインに成功しました")
}
