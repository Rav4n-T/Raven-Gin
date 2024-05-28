package app

import (
	"github.com/gin-gonic/gin"

	"Raven-gin/app/common/request"
	"Raven-gin/app/common/response"
	"Raven-gin/app/services"
)

func Register(c *gin.Context) {
	var form request.Register

	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}

}

func RegisterWithMobile(c *gin.Context) {
	var form request.RegisterWithMobile

	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.RegisterWithMobile(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		response.Success(c, user)
	}

}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.FailByNotFound(c, err.Error())
		return
	}

	response.Success(c, user)
}
