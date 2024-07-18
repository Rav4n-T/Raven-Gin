package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"Raven-Admin/app/common/request"
	"Raven-Admin/app/common/response"
	"Raven-Admin/app/services"
	g "Raven-Admin/global"
	"Raven-Admin/utils"
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

func Logout(c *gin.Context) {
	err := utils.JoinBlacklist(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.Fail(c, "注销失败")
		return
	}

	response.Success(c, nil)
}
