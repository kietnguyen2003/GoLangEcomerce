package controllers

import (
	"net/http"
	"src/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) Users(c *gin.Context) {
	name := c.DefaultQuery("name", "kietnguyen2003")
	c.JSON(http.StatusOK, gin.H{
		"name: ":  name,
		"users: ": uc.userService.Users(),
	})
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	name := c.DefaultQuery("name", "kietnguyen2003")
	c.JSON(http.StatusOK, gin.H{
		"name: ":  name,
		"users: ": uc.userService.GetUserById(userId),
	})
}
