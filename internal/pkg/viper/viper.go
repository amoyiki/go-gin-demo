package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"go-gin-demo/internal/global"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		config = "configs/config.yaml"
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值，config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configs file: %s \n", err))
	}
	v.WatchConfig()
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v

}
