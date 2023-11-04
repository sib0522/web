package models

import (
	"GoEcho/app/domain/repo"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

func Login(c echo.Context) error {
	ok, _ := IsLogin(c)
	if ok {
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

	err = SessionLogin(c, adminAccount.Email())
	if err != nil {
		return xerrors.Errorf("err: %w", err)
	}

	return c.String(http.StatusOK, "Login successful!")
}
