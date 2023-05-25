package mysql

import (
	"fmt"
	"gin-web-scaffolding/models"
	"gin-web-scaffolding/settings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @file      : mysql.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// 声明一个私有的db
var db *gorm.DB

// Init 初始化数据库连接
func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Connect mysql failed", zap.Error(err))
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("db.DB() failed", zap.Error(err))
		return
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	return
}

// Migrate  创建表结构的方法
func Migrate() error {
	err := db.AutoMigrate(&models.User{})
	return err
}
