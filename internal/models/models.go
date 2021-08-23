package models

import (
	"fmt"
	"go-gin-demo/internal/global"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	// Created time.Time `json:"created"`
	// Updated time.Time `json:"updated"`
	Status string `json:"status"`
}

func NewDBEngine() *gorm.DB {
	mysqlconf := global.CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlconf.Username,
		mysqlconf.Password,
		mysqlconf.Host,
		mysqlconf.Port,
		mysqlconf.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), gormConfig())
	fmt.Println(dsn)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.CONFIG.Mysql.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                            // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		}}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)
	config.Logger = newLogger
	return config
}
