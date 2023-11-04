package models

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("secret-key")))

type SessionStore struct {
	Store *sessions.CookieStore
}

func NewSessionStore() *SessionStore {
	return &SessionStore{Store: sessions.NewCookieStore([]byte("secret-key"))}
}

// ログインしているかチェック
func IsLogin(c echo.Context) (bool, string) {
	session, _ := Store.Get(c.Request(), "session")
	if session.Values["email"] == nil || session.Values["email"] == "" {
		return false, ""
	}
	return true, session.Values["email"].(string)
}

// ログイン状態にする
func SessionLogin(c echo.Context, email string) error {
	// ログイン成功時にセッションにユーザーIDを保存
	session, _ := Store.Get(c.Request(), "session")
	session.Values["email"] = email

	// セッションの有効期限を設定（秒単位）
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   120, // 60秒
		Secure:   false,
		HttpOnly: true,
	}

	err := session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return nil
}

// ログアウトする
func SessionLogout(c echo.Context) error {
	// セッション取得
	session, _ := Store.Get(c.Request(), "session")

	session.Values["email"] = ""
	Common.LoginEmail = ""
	Common.IsLogin = false
	err := session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return nil
}
