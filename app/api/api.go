package api

import (
	"GoEcho/app/api/ApiUserLogin"
	"GoEcho/app/api/ApiUserObatain"
	"GoEcho/app/application"
	"GoEcho/app/domain/repo"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	api := c.Request().RequestURI
	switch api {
	case "/user/obatain":
		req := &ApiUserObatain.Request{}
		Unmarshal(c, req)

		userObatainService := application.NewUserObatainService(
			repo.NewUserStatusRepo(),
		)

		res, err := userObatainService.UserObatainService(req)
		if err != nil {
			return err
		}

		return Marshal(c, res)

	case "/user/login":
		req := &ApiUserLogin.Request{}
		Unmarshal(c, req)

		userLoginService := application.NewUserLoginService(
			repo.NewUserStatusRepo(),
		)

		res, err := userLoginService.UserLoginService(req)
		if err != nil {
			return err
		}

		return Marshal(c, res)
	}
	return nil
}
