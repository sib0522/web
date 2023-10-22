package main

import (
	"GoEcho/router"
	"GoEcho/web/models"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware.Loggerブラウザで行った動作がログで表示される
	e.Use(middleware.Logger())

	// 通信を行うたびにログイン状態を確認
	e.Use(checkSession)

	fmt.Println("____________[Init]Router____________")
	router := router.Router{}
	router.InitRouting(e)

	fmt.Println("____________[Init]Server____________")
	e.Logger.Fatal(e.Start(":1323"))

	//go:generate go run github.com/shamaton/msgpackgen
}

func checkSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// セッション取得(なければ生成)
		session, _ := models.Store.Get(c.Request(), "session") // 2

		// ログイン中のアカウントを取得
		val := session.Values["email"]

		if val != nil && val != "" {
			models.Common.LoginEmail = val.(string)
			models.Common.IsLogin = true
		} else {
			models.Common.IsLogin = false
		}

		return next(c)
	}
}
