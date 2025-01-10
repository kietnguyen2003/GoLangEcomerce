package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResonseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccesResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResonseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, ResonseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
