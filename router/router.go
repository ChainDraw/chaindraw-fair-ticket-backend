package router

import (
	v1 "chaindraw-fair-ticket-backend/api/v1"
	"chaindraw-fair-ticket-backend/api/v1/chaindraw"
	"chaindraw-fair-ticket-backend/api/v1/concert"
	"chaindraw-fair-ticket-backend/api/v1/user"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "chaindraw-fair-ticket-backend/docs" // 导入自动生成的文档
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	allowedOrigins := []string{"http://localhost:3000", "http://chaindraw.biturd.com", "https://chaindraw.biturd.com"}
	r.Use(v1.CORSMiddleware(allowedOrigins))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "404 PAGE NOT FOUND!",
		})
	})

	// 设置会话选项
	user.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 一周
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true, // 设置为 true 时，仅允许在 HTTPS 连接中使用
		Domain: ".biturd.com",
	}

	// 集成 Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		//chaindrawApiGroup.POST("lottery_record", chaindraw.LotteryRecordAdd)
	}

	// 演唱会
	concertApiGroup := apiGroup.Group("concert")
	{
		concertApiGroup.GET("concert_list", chaindraw.ConcertList)
		concertApiGroup.POST("commit", concert.ConcertAdd)                //7. 演唱会主办方提交信息
		concertApiGroup.GET("update_status", concert.ConcertStatusUpdate) //7. 演唱会主办方提交信息
		concertApiGroup.POST("review", chaindraw.ReviewConcert)           //8. 演唱会信息审核
		concertApiGroup.POST("cancel", chaindraw.CancelConcert)           //9. 演唱会取消及门票、抵押品退回
	}

	// 用户逻辑相关 路由组
	userApiGroup := apiGroup.Group("user")
	{
		//userApiGroup.POST("login", user.Login)
		//userApiGroup.POST("register", user.Register)
		userApiGroup.GET("nonce", user.Nonce)
		userApiGroup.POST("verify", user.Verify)
		userApiGroup.GET("personal_information", user.PersonalInformation)
		userApiGroup.GET("logout", user.Logout)

		userApiGroup.GET("session_test", user.SessionDemo)

	}

	// 抽奖相关 路由组
	lottery := apiGroup.Group("lottery")
	{
		lottery.GET("list", chaindraw.LotteryListGet)
		lottery.GET("ticketList", chaindraw.TicketListGet)
	}

	// event相关 路由组
	apiGroup.Group("event")
	{
	}

	return r
}
