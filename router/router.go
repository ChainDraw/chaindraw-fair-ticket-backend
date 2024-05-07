package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "404 PAGE NOT FOUND!",
		})
	})

	apiGroup := r.Group("/api/v1")

	// consoleApiGroup 后台路由组
	consoleApiGroup := apiGroup.Group("console")
	{
		consoleApiGroup.POST("login")
	}

	// chaindrawApiGroup 前台路由组
	chaindrawApiGroup := apiGroup.Group("chaindraw")
	{
		chaindrawApiGroup.GET("ticket")
	}

	return r
}
