package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"Raven-gin/app/controllers/app"
	"Raven-gin/app/middleware"
	g "Raven-gin/global"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/register", app.Register)
	router.POST("/registerwithmobile", app.RegisterWithMobile)

	router.POST("/login", app.Login)
	router.POST("/loginwithmobile", app.LoginWithMobile)

	authRouter := router.Group("/auth").Use(middleware.JWTAuth(g.Cof.App.AppName))
	{
		authRouter.POST("/info", app.Info)
	}
}
