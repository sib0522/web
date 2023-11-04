package controllers

import (
	"GoEcho/web/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	renderMap := map[string]any{
		"Common": models.Common,
	}
	return c.Render(http.StatusOK, "index", &renderMap)
}

func User(c echo.Context) error {
	renderMap := map[string]any{
		"Common": models.Common,
	}
	return c.Render(http.StatusOK, "user", &renderMap)
}

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func Resource(c echo.Context) error {
	renderMap := map[string]any{
		"Common": models.Common,
	}
	return c.Render(http.StatusOK, "file_upload", &renderMap)
}

func Gallery(c echo.Context) error {
	gallery, err := models.CreateGallery()
	if err != nil {
		return err
	}
	renderMap := map[string]any{
		"Common":  models.Common,
		"Gallery": gallery,
	}
	return c.Render(http.StatusOK, "gallery", &renderMap)
}

func Delete(c echo.Context) error {
	return c.Render(http.StatusOK, "delete", nil)
}

func Register(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func Password(c echo.Context) error {
	return c.Render(http.StatusOK, "password", nil)
}

func TableList(c echo.Context) error {
	renderMap := map[string]any{
		"Common": models.Common,
	}
	if c.Param("name") == "game" {
		return c.Render(http.StatusOK, "table_list_game", &renderMap)
	} else {
		return c.Render(http.StatusOK, "table_list_admin", &renderMap)
	}
}

func TableDetail(c echo.Context) error {
	table, err := models.CreateTable(c.Param("name"))
	if err != nil {
		return err
	}
	renderMap := map[string]any{
		"Common": models.Common,
		"Table":  table,
	}
	return c.Render(http.StatusOK, "table_detail", &renderMap)
}
