package controller

// @file      : response_code.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// ResponseCode 定义一个类型
type ResponseCode int64

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParam
	CodeServerBusy
	CodeRateLimit
	CodeNeedLogin
	CodeInvalidToken
	CodeUserExist
	CodeUserNotExists
)

var CodeMsgMap = map[ResponseCode]string{
	CodeSuccess:       "success",
	CodeInvalidParam:  "请求参数错误",
	CodeServerBusy:    "服务繁忙",
	CodeRateLimit:     "访问请求被限制",
	CodeNeedLogin:     "需要登录",
	CodeInvalidToken:  "无效的Token",
	CodeUserExist:     "用户已存在",
	CodeUserNotExists: "用户尚未注册",
}

func (c ResponseCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
