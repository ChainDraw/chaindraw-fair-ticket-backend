package global

import (
	"chaindraw-fair-ticket-backend/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	APP_CONFIG *config.AppConfig
	LOGGER     *zap.Logger
	DB         *gorm.DB
)
