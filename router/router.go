package router

import (
	"GoEcho/controllers"
	"GoEcho/forms"
	"GoEcho/logic"
	"fmt"
	"github.com/labstack/echo/v4"
)

func InitRouting(e *echo.Echo) {
	fmt.Println("____________[Init]Router____________")

	// 静的ページの読み込み設定
	e.Static("/public", "public")
	forms.SetRenderer(e, "")

	InitPages(e)
	InitLogics(e)
}

// InitPages 各ページのルーター処理を設定する
func InitPages(e *echo.Echo) {
	/*
		jwtConfig := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(logic.TokenData)
			},
			SigningKey: []byte("jwtUserLogin"),
		}
	*/

	e.GET("/", controllers.Index)
	e.GET("/account/register", controllers.Register)
	e.GET("/account/login", controllers.Login)
	e.GET("/account/password", controllers.Password)
	e.GET("/account/delete", controllers.Delete)
	e.GET("/file/upload", controllers.FileUpload)
	e.GET("/file/gallery", controllers.Gallery)
	e.GET("/table/admin", controllers.TableList)
	e.GET("/table/admin/:name", controllers.TableDetail)
	//e.GET("/table/game", controllers.TableList)
	//e.GET("/table/game/:name", controllers.TableDetail)
	e.GET("/ws", controllers.Chat)
	e.GET("/layout/color", controllers.Layout)

	/*
		tableGroup := e.Group("/table")
		tableGroup.Use(echojwt.WithConfig(jwtConfig))
		//tableGroup.GET("", logic.JWTCheck)
		tableGroup.GET("/admin", controllers.TableList)
		tableGroup.GET("/game", controllers.TableList)
	*/
}

// InitLogics Logic周りのルーター処理を設定する
func InitLogics(e *echo.Echo) {
	e.POST("/logic/signup", logic.SignUp)
	e.POST("/account/login", logic.Login)
	e.POST("/logic/sendmail", logic.SendMail)
	e.POST("/logic/delete", logic.Delete)
	e.GET("/logic/logout", logic.Logout)
	e.POST("/upload/images", logic.UploadImage)
}
