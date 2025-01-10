package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ErrorHandleMiddleware")
		c.Next()
	}
}
