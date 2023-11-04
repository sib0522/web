package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vmihailenco/msgpack/v5"
)

type IRequest interface {
	Unmarshal(c echo.Context)
}

type IResponse interface {
	Marshal(c echo.Context)
}

func Unmarshal(c echo.Context, r IRequest) error {
	body, err := c.Request().GetBody()
	if err != nil {
		return err
	}

	err = msgpack.NewDecoder(body).Decode(r)
	if err != nil {
		return err
	}

	return nil
}

func Marshal(c echo.Context, r IResponse) error {
	msgRes, err := msgpack.Marshal(r)
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationMsgpack, msgRes)
}
