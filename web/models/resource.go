package models

import (
	"GoEcho/web/lib"
	"net/http"

	"github.com/labstack/echo/v4"
)

// リソースをアップロードする
func UploadResource(c echo.Context) error {
	multiFile, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := multiFile.File["files[]"]

	aws := lib.NewAWSService()
	aws.UploadMultiple(files)

	return c.String(http.StatusOK, "File %s uploaded successfully.")
}

// リソースをダウンロードする
func DownloadResource(c echo.Context) error {
	aws := lib.NewAWSService()
	if err := aws.DownloadMultiple("uploads/"); err != nil {
		return err
	}
	return nil
}
