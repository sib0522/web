package models

import (
	"GoEcho/app/domain/repo"
	"GoEcho/app/util/clock"
	"GoEcho/app/util/convert"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserStatusData struct {
	Uuid      string `json:"uuid"`
	Level     uint32 `json:"level"`
	Exp       uint64 `json:"exp"`
	Money     uint64 `json:"money"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

func CheckUserId(c echo.Context) error {
	idStr := c.FormValue("userId")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return err
	}

	userStatus, err := repo.NewUserStatusRepo().ReadById(uint32(id))
	if err != nil {
		return err
	}

	userStatusData := &UserStatusData{
		Uuid:      userStatus.Uuid(),
		Level:     userStatus.Level(),
		Exp:       userStatus.Exp(),
		Money:     userStatus.Money(),
		UpdatedAt: userStatus.UpdatedAt().Format(clock.DateTimeFormat),
		CreatedAt: userStatus.CreatedAt().Format(clock.DateTimeFormat),
	}

	return c.JSON(http.StatusOK, userStatusData)
}

func UpdateUserData(c echo.Context) error {
	target := c.Param("name")
	v := c.FormValue("id")

	userId, err := convert.ToUint32(v)
	if err != nil {
		return err
	}

	userStatus, err := repo.NewUserStatusRepo().ReadById(userId)
	if err != nil {
		return err
	}

	v = c.FormValue("value")
	value, err := convert.ToUint64(v)
	if err != nil {
		return err
	}

	e := userStatus.Entity()
	switch target {
	case "level":
		e.Level = uint32(value)
	case "exp":
		e.Exp = value
	case "money":
		e.Money = value
	}
	e.Apply(userStatus)

	err = repo.NewUserStatusRepo().CreateOrUpdateByModel(userStatus)
	if err != nil {
		return err
	}
	return nil
}
