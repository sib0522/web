package main

import (
	"GoEcho/database"
	"GoEcho/router"
	"GoEcho/web/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()

	// middleware.Loggerブラウザで行った動作がログで表示される
	e.Use(middleware.Logger())

	// 通信を行うたびにログイン状態を確認
	e.Use(checkSession)

	// Router初期化
	fmt.Println("____________[Init]Router____________")
	router.Initialize(e)

	// DB接続
	fmt.Println("____________[Init]DB________________")
	database.ConnectDB()

	// サーバー起動
	fmt.Println("____________[Init]Server____________")

	err := godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("DOMAIN_ADMIN_DEV")),
		Cache:      autocert.DirCache("certs"),
	}

	e.Listener = certManager.Listener()

	e.Logger.Fatal(e.StartTLS(":1323", "", ""))
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
