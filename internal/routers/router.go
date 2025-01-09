package routers

import (
	// thấy controller dài qua thì để chữ c vô
	c "src/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/api")
	{
		// Pong
		v1.GET("/ping/", c.NewPongController().Pong)
		v1.GET("/ping/:name", c.NewPongController().Pong)

		// Users
		v1.GET("/users", c.NewUserController().Users)
		v1.GET("/users/:id", c.NewUserController().GetUserById)
	}
	return r
}
