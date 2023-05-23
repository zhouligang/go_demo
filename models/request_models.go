package models

// @file      : request_models.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SelectUser 查询相关用户信息的结构体
type SelectUser struct {
	UserID       int64
	Username     string
	Password     string
	AccessToken  string
	RefreshToken string
}
