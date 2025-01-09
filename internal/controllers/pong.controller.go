package controllers

import (
	"net/http"
	"src/internal/service"

	"github.com/gin-gonic/gin"
)

type PongController struct {
	pongService *service.PongService
} // định nghĩa một struct PongController

func NewPongController() *PongController { // tạo một hàm NewPongController trả về một con trỏ đến PongController
	return &PongController{
		pongService: service.NewPongService(),
	}
}

// Tạo method Pong cho PongController
func (pc *PongController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + pc.pongService.Pong(),
	})
}
