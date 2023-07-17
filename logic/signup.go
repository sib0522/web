package logic

import (
	"GoEcho/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c echo.Context) error {
	nickName := c.FormValue("name")
	email := c.FormValue("email")

	p := c.FormValue("password")
	confirm := c.FormValue("password2")

	if nickName == "" || email == "" || p == "" || confirm == "" {
		return c.Redirect(http.StatusFound, "/register")
		/*
			return c.Render(http.StatusOK, "/register", map[string]bool{
				"isShowAlert": true,
			})
		*/
	}

	if p != confirm {
		return c.String(http.StatusExpectationFailed, "入力したパスワードを確認してください")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 2)

	password := string(hash)

	account := models.Account{
		Nickname: nickName,
		Password: password,
		Email:    email,
	}

	if b := account.Read(); b == true {
		return c.String(http.StatusExpectationFailed, "既に存在しているアカウント名です")
	}

	account.Create()

	return c.Render(http.StatusOK, "success", nil)
}
