package redis

import (
	"context"
	"fmt"
	"go-gin-demo/global"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port),
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, canal := context.WithTimeout(context.Background(), 5*time.Second)
	defer canal()

	_, err := client.Ping().Result()
	if err != nil {
		global.LOG.Error("redis connect ping fail, err:", zap.Any("err", err))
	}
	global.REDIS = client
}
