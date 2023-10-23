package models

import (
	"GoEcho/app/domain/repo"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

type TokenData struct {
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	session, _ := Store.Get(c.Request(), "session")
	fmt.Println(session.Values["email"])
	if session.Values["email"] != nil && session.Values["email"] != "" {
		fmt.Println("既にログインしています")
		return c.Redirect(404, "/")
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	adminAccountRepo := repo.NewAdminAccountRepo()
	adminAccount, err := adminAccountRepo.ReadByEmail(email)
	if err != nil {
		return c.String(404, "ログインに失敗しました")
	}
	if adminAccount.Email() == "" {
		return xerrors.Errorf("アカウントが存在しません")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(adminAccount.Password()), []byte(password)); err != nil {
		return c.String(404, "パスワードが一致しません")
	}

	// ログイン成功時にセッションにユーザーIDを保存
	session, _ = Store.Get(c.Request(), "session")
	session.Values["email"] = adminAccount.Email() // ここに実際のユーザーIDをセット

	// セッションの有効期限を設定（秒単位）
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   120, // 60秒
		Secure:   false,
		HttpOnly: true,
	}

	session.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Login successful!")
}
