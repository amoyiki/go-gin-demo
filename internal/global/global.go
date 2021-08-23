package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go-gin-demo/configs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG configs.Server
	DB     *gorm.DB
	LOG    *zap.Logger
	CFG    *viper.Viper
	REDIS  *redis.Client
)
