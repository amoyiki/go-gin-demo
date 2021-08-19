package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go-gin-demo/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	LOG    *zap.Logger
	CFG    *viper.Viper
	REDIS  *redis.Client
)
