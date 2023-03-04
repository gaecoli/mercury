package app

import (
	"github.com/gin-gonic/gin"
	"mercury/mercury/app/common/request"
	"mercury/mercury/app/common/response"
	"mercury/mercury/app/services"
)

func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFailed(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFailed(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
