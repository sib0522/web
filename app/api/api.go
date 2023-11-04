package api

import (
	"GoEcho/app/api/ApiUserLogin"
	"GoEcho/app/api/ApiUserObatain"
	"GoEcho/app/application"
	"GoEcho/app/domain/repo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack/v5"
)

func Handler(c echo.Context) error {
	api := c.Request().RequestURI
	switch api {
	case "/user/obatain":
		body, err := c.Request().GetBody()
		if err != nil {
			return err
		}

		request := &ApiUserObatain.Request{}
		err = msgpack.NewDecoder(body).Decode(request)
		if err != nil {
			return err
		}

		userObatainService := application.NewUserObatainService(
			repo.NewUserStatusRepo(),
		)

		res, err := userObatainService.UserObatainService(request)
		if err != nil {
			return err
		}

		msgRes, err := msgpack.Marshal(res)
		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, echo.MIMEApplicationMsgpack, msgRes)

	case "/user/login":
		/*
			body, err := c.Request().GetBody()
			if err != nil {
				return err
			}

			request := &ApiUserLogin.Request{}
			err = msgpack.NewDecoder(body).Decode(request)
			if err != nil {
				return err
			}

			userLoginService := application.NewUserLoginService(
				repo.NewUserStatusRepo(),
			)

			res, err := userLoginService.UserLoginService(request)
			if err != nil {
				return err
			}

			msgRes, err := msgpack.Marshal(res)
			if err != nil {
				return err
			}
		*/

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
