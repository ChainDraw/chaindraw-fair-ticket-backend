package main

import (
	"chaindraw-fair-ticket-backend/core"
	"chaindraw-fair-ticket-backend/global"
	"chaindraw-fair-ticket-backend/go2chain"
	"chaindraw-fair-ticket-backend/initialize"
)

// @title Your Project API
// @version 1.0
// @description This is a sample API for Your Project using Gin and Swagger
// @BasePath /api/v1
func main() {
	// 加载配置文件
	core.Viper()

	// 加载日志库
	global.LOGGER = core.Zap()

	// 初始化 MySQL 数据库
	global.DB = initialize.Gorm()

	// 初始化 Redis 数据库

	//  初始化chain监听
	go2chain.ListerInit(global.DB)
	go go2chain.Run()

	// 启动服务
	core.RunServer()

}
