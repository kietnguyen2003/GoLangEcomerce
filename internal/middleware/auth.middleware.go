package middleware

import (
	"fmt"
	"src/package/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			fmt.Println("Unauthorized")
			response.ErrorResponse(c, response.ErrUnauthorized, "")
			c.Abort()
			return
		}
		fmt.Println("AuthMiddleware")
		c.Next()

	}
}
