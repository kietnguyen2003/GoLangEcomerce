package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CrosMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CrosMiddleware")
		c.Next()
	}
}
