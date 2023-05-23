package logic

import (
	"gin-web-scaffolding/dao/mysql"
	"gin-web-scaffolding/models"
	"gin-web-scaffolding/utils"
)

// @file      : user.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// SignUp 处理用户登录
func SignUp(pUser *models.ParamsSignUp) (err error) {
	// 判断用户是否已经存在
	if err = mysql.CheckUserExists(pUser.Username); err != nil {
		return
	}

	// 生成UID
	userID := utils.GenSnowFlakeID()

	// 构造一个User实例
	u := models.User{
		UserID:   userID,
		Username: pUser.Username,
		Password: pUser.Password,
		Email:    pUser.Email,
	}

	// 写入数据库
	err = mysql.InsertUser(&u)
	return
}

// Login 处理用户登录
func Login(p *models.ParamsLogin) (user *models.SelectUser, err error) {
	user = &models.SelectUser{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	// 生成JWT token
	accessToken, refreshToken, err := utils.GenJWTToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
