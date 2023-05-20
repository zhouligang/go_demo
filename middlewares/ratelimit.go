package middlewares

import (
	"gin-web-scaffolding/controller"
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

// @file      : ratelimit.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// RateLimitMiddleware 基于令牌桶的限制流量方案
// 参数代表的意思是：每fillInterval的时间新增加cap个令牌
func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(context *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			context.String(http.StatusOK, controller.CodeRateLimit.Msg())
			context.Abort()
			return
		}
		// 取到了令牌就放行
		context.Next()

	}
}
