package logic

import (
	"GoEcho/database"
	"GoEcho/models"
	"fmt"
	"github.com/labstack/echo/v4"
)

func CreateTable(c echo.Context) error {
	db := database.ConnectDB()
	if db == nil {
		return nil
	}

	var user models.User
	columnNames := user.ReadColumns()

	var cols string

	for _, col := range columnNames {
		cols = cols + fmt.Sprintf("\t\t<th>%v</th>\n\t\t", col)
	}

	return nil
}
