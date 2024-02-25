package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func baseSuccess(c *gin.Context, code int, data interface{}) {

	c.JSON(code, Response{data, "success"})
}

func Success(c *gin.Context, data interface{}) {
	baseSuccess(c, http.StatusOK, data)
}

func baseFail(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{nil, msg})
}

func Fail(c *gin.Context, msg string) {
	baseFail(c, http.StatusBadRequest, msg)
}

func FailByAuth(c *gin.Context, msg string, code int) {
	if code == 0 {
		baseFail(c, http.StatusUnauthorized, msg)
		return
	}
	baseFail(c, http.StatusForbidden, msg)

}

func FailByNotFound(c *gin.Context, msg string) {
	baseFail(c, http.StatusNotFound, msg)
}
