package logic

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadImage(c echo.Context) error {
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

	dst, err := os.Create(filePath)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.", hashed))
}
