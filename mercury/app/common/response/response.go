package response

import (
	"github.com/gin-gonic/gin"
	"mercury/mercury/global"
	"net/http"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

func Failed(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

func FailedByError(c *gin.Context, error global.CustomError) {
	Failed(c, error.ErrorCode, error.ErrorMsg)
}

func ValidateFailed(c *gin.Context, msg string) {
	Failed(c, global.Errors.ValidateArgsError.ErrorCode, msg)
}

func BusinessFailed(c *gin.Context, msg string) {
	Failed(c, global.Errors.NotImplementedError.ErrorCode, msg)
}
