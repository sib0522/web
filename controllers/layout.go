package controllers

import (
	"GoEcho/logic"
	"github.com/labstack/echo/v4"
)

func Layout(c echo.Context) error {
	if c.Path() == "/layout/color" {
		logic.Color(&c)
	}

	return nil
}
