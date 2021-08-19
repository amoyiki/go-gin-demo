package main

import (
	"fmt"
	"go-gin-demo/global"
	"go-gin-demo/models"
	"go-gin-demo/pkg/redis"
	"go-gin-demo/pkg/viper"
	"go-gin-demo/pkg/zap"
	"go-gin-demo/routers"
	"go-gin-demo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	global.CFG = viper.Viper()
	global.LOG = zap.Zap()
	global.DB = models.NewDBEngine()
	fmt.Println(global.CONFIG.System.Port)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err %v\n", err)
		return
	}
	global.DB.Debug()
	redis.Redis()
	gin.SetMode("debug")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.CONFIG.System.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Millisecond,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {

		fmt.Printf("s.ListenAndServer err: %v", err)
	}
}
