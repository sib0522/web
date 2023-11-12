package models

import (
	"github.com/joho/godotenv"
	"net/http"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"
)

func SendMail(c echo.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", os.Getenv("MAIL_ADDRESS"), os.Getenv("MAIL_PASSWORD"), "smtp.gmail.com")
	from := os.Getenv("MAIL_ADDRESS")
	to := []string{os.Getenv("MAIL_ADDRESS")}

	header := "テストメールです\r\n"
	mailType := "\r\n"
	body := "これはテストメールです"
	msg := []byte(header + mailType + body)

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "メールを送りました")
}
