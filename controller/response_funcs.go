package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @file      : response_funcs.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// ResponseData 定义返回给前端的结构体
type ResponseData struct {
	Code ResponseCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data,omitempty"` //omitempty:如果为空则忽略，不返回给前端
}

// ResponseError 定义返回错误的方法
func ResponseError(context *gin.Context, code ResponseCode) {
	context.JSON(http.StatusOK,
		&ResponseData{
			Code: code,
			Msg:  code.Msg(),
			Data: nil,
		})
}

// ResponseErrorWithMsg 定义返回指定错误信息的方法
func ResponseErrorWithMsg(context *gin.Context, code ResponseCode, msg interface{}) {
	context.JSON(http.StatusOK,
		&ResponseData{
			Code: code,
			Msg:  msg,
			Data: nil,
		})
}

// ResponseSuccess  返回成功信息的方法
func ResponseSuccess(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK,
		&ResponseData{
			Code: CodeSuccess,
			Msg:  CodeSuccess.Msg(),
			Data: data,
		})
}
