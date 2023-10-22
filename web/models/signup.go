package models

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/domain/repo"
	"GoEcho/app/util/clock"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

func SignUp(c echo.Context) error {
	nickName := c.FormValue("name")
	email := c.FormValue("email")

	p := c.FormValue("password")
	confirm := c.FormValue("password2")

	if nickName == "" || email == "" || p == "" || confirm == "" {
		return xerrors.Errorf("全て入力してください")
	}

	if p != confirm {
		return c.String(http.StatusExpectationFailed, "入力したパスワードを確認してください")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 2)

	password := string(hash)

	adminAccountRepo := repo.NewAdminAccountRepo()
	adminAccount, err := adminAccountRepo.ReadByEmail(email)
	if err != nil {
	}
	if adminAccount.Id() != 0 {
		return c.String(http.StatusExpectationFailed, "既に存在しているアカウント名です")
	}

	t := clock.Now().Time

	adminAccountRepo.CreateByModel(model.NewAdminAccount(email, password, nickName, t))

	return c.Render(http.StatusOK, "success", nil)
}
