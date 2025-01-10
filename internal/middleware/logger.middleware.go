package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("LoggerMiddleware")
		c.Next()
	}
}
