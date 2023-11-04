package models

import (
	"GoEcho/app/domain/repo"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	ok, email := IsLogin(c)
	if !ok {
		// ログインしてない場合は弾く
		return c.Redirect(http.StatusForbidden, "/")
	}

	adminAccount, err := repo.NewAdminAccountRepo().ReadByEmail(email)
	if err != nil {
		return err
	}

	err = repo.NewAdminAccountRepo().DeleteByModel(adminAccount)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "アカウントを削除しました")
}
