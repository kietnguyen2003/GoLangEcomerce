package initialize

import (
	// thấy controller dài qua thì để chữ c vô

	"src/global"
	"src/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	r.Use() // loggin
	r.Use() // cross
	r.Use() // limiter global

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainRouterGroup := r.Group("/v1/api")
	{
		MainRouterGroup.GET("/checkStatus") //tracking monitor
	}
	{
		manageRouter.AdminRouter.InitAdminRouter(MainRouterGroup)
		manageRouter.UserRouter.InitUserRouter(MainRouterGroup)
	}
	{
		userRouter.UserRouter.InitUserRouter(MainRouterGroup)
		userRouter.ProductRouter.InitProductRouter(MainRouterGroup)
	}
	return r
}
