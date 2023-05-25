package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"gin-web-scaffolding/models"
)

// @file      : user.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

const secret = "八宝糖" //盐，用于对密码加密

// CheckUserExists 判断用户是否存在
func CheckUserExists(username string) error {
	var user models.User
	result := db.First(&user, "username=?", username)
	if result.RowsAffected != 0 {
		return ErrorUserExist
	}
	return nil
}

// encryptPassword 加密密码的函数
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// InsertUser 往数据库中创建增加用户记录的方法
func InsertUser(user *models.User) error {
	user.Password = encryptPassword(user.Password)
	result := db.Create(user)
	return result.Error
}

// Login 处理用户登录的方法，
func Login(user *models.SelectUser) (err error) {
	u := &models.User{}
	result := db.Where("username=? and password=?", user.Username, encryptPassword(user.Password)).Find(u)
	if result.RowsAffected == 0 {
		return ErrorUserNotExists
	} else if result.RowsAffected != 1 {
		return result.Error
	} else {
		user.UserID = u.UserID
	}
	return
}
