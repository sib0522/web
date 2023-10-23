package models

import (
	"GoEcho/web/lib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadResource(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	src, openErr := file.Open()
	if openErr != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer src.Close()

	dir := "./public/images"
	os.MkdirAll(dir, 0777)

	ext := filepath.Ext(file.Filename)
	bt := []byte(file.Filename)
	hashed := sha1.Sum(bt)
	hashedName := hex.EncodeToString(hashed[:])

	filePath := fmt.Sprintf("%v/%v%v", dir, hashedName, ext)

	aws := lib.NewAWSService()
	aws.UploadMultiple([]string{filePath})

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.", hashed))
}

func DownloadResource(c echo.Context) error {
	aws := lib.NewAWSService()
	if err := aws.DownloadMultiple("uploads/public/images/"); err != nil {
		return err
	}
	return nil
}
