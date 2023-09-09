package logic

import (
	"GoEcho/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type TokenData struct {
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	// Viewからemailとpasswordを取得する
	email := c.FormValue("email")
	password := c.FormValue("password")

	account := &models.Account{
		Password: password,
		Email:    email,
	}

	// 読み取りできるアカウントのデータがないと一致するデータが存在しないことにする
	b := account.Read()
	if b == false {
		return c.String(404, "ログインに失敗しました")
	}

	claims := &TokenData{
		account.Nickname,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("jwtUserLogin"))
	if err != nil {
		return err
	}

	c.Request().Header.Add("user", t)

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
