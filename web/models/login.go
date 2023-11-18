package models

import (
	"GoEcho/app/domain/repo"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	ok, _ := IsLogin(c)
	if ok {
		fmt.Println("不正なアクセス")
		return c.String(http.StatusBadRequest, "既にログインしています")
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	adminAccountRepo := repo.NewAdminAccountRepo()
	adminAccount, err := adminAccountRepo.ReadByEmail(email)
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "")
	}
	if adminAccount.Email() == "" {
		return c.String(http.StatusMethodNotAllowed, "アカウントが存在しません")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(adminAccount.Password()), []byte(password)); err != nil {
		return c.String(http.StatusExpectationFailed, "パスワードが一致しません")
	}

	err = SessionLogin(c, adminAccount.Email())
	if err != nil {
		return c.String(http.StatusExpectationFailed, "セッション生成に失敗しました")
	}

	return c.String(http.StatusOK, "ログインしました")
}
