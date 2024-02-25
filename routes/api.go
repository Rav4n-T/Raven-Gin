package routes

import (
	"Raven-gin/app/controllers/app"
	"Raven-gin/app/middleware"
	g "Raven-gin/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/register", app.Register)

	router.POST("/login", app.Login)

	authRouter := router.Group("/auth").Use(middleware.JWTAuth(g.Cof.App.AppName))
	{
		authRouter.POST("/info", app.Info)
	}
}
