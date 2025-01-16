package global

import (
	"src/package/logger"
	"src/package/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Redis  *redis.Client
	Mdb    *gorm.DB
)
