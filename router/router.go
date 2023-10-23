package router

import (
	"GoEcho/app/api"
	"GoEcho/forms"
	"GoEcho/web/controllers"
	"GoEcho/web/models"

	"github.com/labstack/echo/v4"
)

type IRouter interface {
	InitRouting(e *echo.Echo)
}
type Router struct {
	IRouter
}

func (Router) InitRouting(e *echo.Echo) {
	initWebPage(e)
	initLogic(e)
	initApi(e)
}

// InitPages 各ページのルーター処理を設定する
func initWebPage(e *echo.Echo) {
	// 静的ページの読み込み設定
	e.Static("/web/public", "/web/public")
	forms.InitRenderer(e)

	e.GET("/", controllers.Index)
	e.GET("/user", controllers.User)
	e.GET("/account/register", controllers.Register)
	e.GET("/account/login", controllers.Login)
	e.GET("/account/password", controllers.Password)
	e.GET("/account/delete", controllers.Delete)
	e.GET("/file/resource", controllers.Resource)
	e.GET("/file/gallery", controllers.Gallery)
	e.GET("/table/:name", controllers.TableList)
	e.GET("/table/admin/:name", controllers.TableDetail)
}

// InitLogics Logic周りのルーター処理を設定する
func initLogic(e *echo.Echo) {
	e.POST("/account/register", models.SignUp)
	e.POST("/account/login", models.Login)
	e.POST("/logic/sendmail", models.SendMail)
	e.POST("/logic/delete", models.Delete)
	e.GET("/account/logout", models.Logout)
	e.POST("/resource/upload", models.UploadResource)
	e.POST("/resource/download", models.DownloadResource)
	e.POST("/user", models.CheckUserId)
}

// ゲームロジックAPIを設定
func initApi(e *echo.Echo) {
	e.POST("/user/:apiName", api.Handler)
}
