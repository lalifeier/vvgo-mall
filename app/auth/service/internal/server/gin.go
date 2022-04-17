package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lalifeier/vvgo-mall/app/auth/service/internal/service"
)

func NewGinServer(authService *service.AuthService) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/*")

	// router.Use(kgin.Middleware())

	api := router.Group("/")
	{
		api.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"title": "login",
			})
		})
		api.POST("/login", authService.Login)
		api.GET("/auth", authService.Auth)

		api.POST("/authorize", authService.Authorize)
		api.POST("/token", authService.Token)
	}

	return router
}
