package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.GET("/:id")
	}
	// private
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
		// userRouterPrivate.GET("/my_infor")
	}
}
