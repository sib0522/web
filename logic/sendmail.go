package logic

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/smtp"
)

func SendMail(c echo.Context) error {
	auth := smtp.PlainAuth("", "icchi522@gmail.com", "pgjmzlvpfrugyxoh", "smtp.gmail.com")
	from := "icchi522@gmail.com"
	to := []string{"icchi522@gmail.com"}

	header := "テストメールです\r\n"
	mailType := "\r\n"
	body := "これはテストメールです"
	msg := []byte(header + mailType + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "メールを送りました")
}
