package app

import (
	"Raven-gin/app/common/request"
	"Raven-gin/app/common/response"
	"Raven-gin/app/services"
	g "Raven-gin/global"
	"Raven-gin/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.Fail(c, err.Error())
	} else {
		tokenData, err, _ := utils.CreateToken(g.Cof.App.AppName, user)
		if err != nil {
			response.Fail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}

}
