package logic

import (
	"GoEcho/constants"
	"github.com/labstack/echo/v4"
)

func Color(c *echo.Context) {
	cp := *c
	result := cp.Get("context")
	var common, _ = result.(constants.Common)
	common.SideNaviColor = "light"
	cp.Set("context", common)
}
