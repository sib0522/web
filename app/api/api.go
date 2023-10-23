package api

import (
	"GoEcho/app/api/ApiUserLogin"
	"GoEcho/app/application"
	"GoEcho/app/domain/repo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack/v5"
)

func Handler(c echo.Context) error {
	api := c.Request().RequestURI
	switch api {
	case "/user/update":

	case "/user/login":
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

		return c.Blob(http.StatusOK, echo.MIMEApplicationMsgpack, msgRes)
	}

	return nil
}
