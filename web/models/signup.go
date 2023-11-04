package models

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/domain/repo"
	"GoEcho/app/util/clock"
	"GoEcho/database"
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
		return xerrors.Errorf("err : %w", "全て入力してください")
	}

	if p != confirm {
		return c.String(http.StatusExpectationFailed, "パスワードを確認してください")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 2)
	password := string(hash)

	adminAccountRepo := repo.NewAdminAccountRepo()
	adminAccount, err := adminAccountRepo.ReadByEmail(email)
	if database.IsErrorNoRows(err) {
		err = nil
	}
	if err != nil {
		return err
	}

	if adminAccount != nil && adminAccount.Id() != 0 {
		return c.String(http.StatusExpectationFailed, "既に存在しているアカウント名です")
	}

	t := clock.Now().Time

	adminAccountRepo.CreateByModel(model.NewAdminAccount(email, password, nickName, t))

	return c.Render(http.StatusOK, "success", nil)
}
