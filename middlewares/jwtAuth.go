package middlewares

import (
	"gin-web-scaffolding/controller"
	"gin-web-scaffolding/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// @file      : jwtAuth.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(context, controller.CodeNeedLogin)
			context.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 3)
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			controller.ResponseError(context, controller.CodeInvalidToken)
			context.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		_, err := utils.ParseJWTToken(parts[1])
		if err != nil {
			parts[1], parts[2], err = utils.RefreshToken(parts[1], parts[2])
			if err != nil {
				controller.ResponseError(context, controller.CodeInvalidToken)
				context.Abort()
				return
			}
		}
		mc, _ := utils.ParseJWTToken(parts[1])

		// 将当前请求的username信息保存到请求的上下文context中
		context.Set(controller.ContextUserIDKey, mc.Username)
		context.Set(controller.ContextAccessToken, parts[1])
		context.Set(controller.ContextRefreshToken, parts[2])
		context.Next()
	}
}
