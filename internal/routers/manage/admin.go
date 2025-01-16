package manage

import (
	"src/internal/controllers"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (arg *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
		adminRouterPublic.GET("/users", controllers.NewUserController().Users)
		adminRouterPublic.GET("/users/:id", controllers.NewUserController().GetUserById)
	}
	// // private
	// adminRouterPrivate := Router.Group("/admin/user")
	// // adminRouterPrivate.Use(limiter())
	// // adminRouterPrivate.Use(Authen())
	// // adminRouterPrivate.Use(Permission())
	// {
	// 	adminRouterPrivate.GET("/my_infor")
	// }

}
