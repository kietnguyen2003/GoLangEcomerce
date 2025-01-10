package controllers

import (
	"src/internal/service"
	"src/package/response"
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
	users := uc.userService.Users()
	response.SuccesResponse(c, response.ErrorCodeSuccess, users)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(c, response.ErrorCodeParamInvalid, "param invalid")
		return
	}
	response.SuccesResponse(c, response.ErrorCodeSuccess, uc.userService.GetUserById(userId))
}
