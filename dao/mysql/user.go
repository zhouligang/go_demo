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

const secret = "八宝糖" //盐

func CheckUserExists(username string) error {
	var user models.User
	result := db.First(&user, "username=?", username)
	if result.RowsAffected != 0 {
		return ErrorUserExist
	}
	return nil
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func InsertUser(user *models.User) error {
	user.Password = encryptPassword(user.Password)
	result := db.Create(user)
	return result.Error
}

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
