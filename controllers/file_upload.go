package controllers

import (
	"GoEcho/forms"
	"github.com/labstack/echo/v4"
	"net/http"
)

func FileUpload(c echo.Context) error {
	forms.SetRenderer(c.Echo(), "views/file_upload.html")
	return c.Render(http.StatusOK, "file_upload", nil)
}
