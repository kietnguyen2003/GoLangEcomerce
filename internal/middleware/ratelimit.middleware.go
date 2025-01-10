package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("RateLimitMiddleware")
		c.Next()
	}
}
