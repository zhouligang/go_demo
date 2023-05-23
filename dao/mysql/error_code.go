package mysql

import "errors"

// @file      : error_code.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExists = errors.New("用户未注册")
)
