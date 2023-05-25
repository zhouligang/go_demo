package models

import (
	"gorm.io/gorm"
)

// @file      : db_models.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// User 创建数据库表对应的model
type User struct {
	gorm.Model
	UserID   int64  `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
	Email    string
	Gender   int8 `gorm:"default:0"`
}
