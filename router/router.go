package router

import (
	"GoEcho/controllers"
	"GoEcho/logic"
	"fmt"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	fmt.Println("____________[Init]Router____________")

	// 静的ページの読み込み設定
	e.Static("/public", "public")

	InitPages(e)
	InitLogics(e)
}

// InitPages 各ページのルーター処理を設定する
func InitPages(e *echo.Echo) {
	e.GET("/", controllers.Index)
	e.GET("/register", controllers.Register)
	e.GET("/login", controllers.Login)
	e.GET("/password", controllers.Password)
	e.GET("/file_upload", controllers.FileUpload)
}

// InitLogics Logic周りのルーター処理を設定する
func InitLogics(e *echo.Echo) {
	e.POST("/logic/signup", logic.SignUp)
	e.POST("/logic/login", logic.Login)
	e.POST("/logic/sendmail", logic.SendMail)
}
