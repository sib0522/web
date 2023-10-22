package models

import (
	"GoEcho/app/domain/repo"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CheckUserId(c echo.Context) error {
	idStr := c.FormValue("userId")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return err
	}

	adminAccount, err := repo.NewAdminAccountRepo().ReadById(uint32(id))
	if err != nil {
		return err
	}
	//return c.String(http.StatusOK, adminAccount.Email())
	return c.Redirect(http.StatusOK, idStr)
}
