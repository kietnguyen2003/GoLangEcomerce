package users

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/login")
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}
	// private
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/my_infor")
	}
}
