package router

import (
	"chaindraw-fair-ticket-backend/api/v1/user"
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

	// 用户逻辑相关 路由组
	userApiGroup := apiGroup.Group("user")
	{
		userApiGroup.POST("login", user.Login)
		userApiGroup.POST("register", user.Register)
	}

	return r
}
