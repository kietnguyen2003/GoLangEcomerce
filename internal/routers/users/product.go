package users

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public
	productRouterPublic := Router.Group("/user")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail")
	}
	// private

}
