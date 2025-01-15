package initialize

import (
	// thấy controller dài qua thì để chữ c vô
	c "src/internal/controllers"
	"src/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default() // dùng crt + click để xem hàm này sử lý Middleware như thế nào

	// Middleware sau khi click vào sẽ thấy
	// engine.Use(Logger(), Recovery()) click tiếp vô Logger(), sẽ thấy middleware đó sử dụng HandleFunc
	// Do đó tất cả các Middleware đều sử dụng HandleFunc, xem các middleware bên dưới để thấy rõ hơn
	r.Use(middleware.CrosMiddleware(), middleware.LoggerMiddleware(), middleware.ErrorHandleMiddleware(), middleware.RateLimitMiddleware(), middleware.AuthMiddleware())

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
