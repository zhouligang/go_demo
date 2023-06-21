package routes

import (
	"gin-web-scaffolding/controller"
	"gin-web-scaffolding/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"

	_ "gin-web-scaffolding/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// @file      : routes.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// SetupRoutes 设置路由
func SetupRoutes(mode string, AllowOrigins, AllowMethods []string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为release模式
	}
	//  配置跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = AllowOrigins
	config.AllowMethods = AllowMethods
	config.AllowHeaders = []string{"Origin"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	router := gin.New()
	//  使用自定义的三个中间件
	router.Use(
		cors.New(config),
		middlewares.GinLogger(),
		middlewares.GinRecovery(true),
		middlewares.RateLimitMiddleware(2, 5),
	)

	//  定义路由组
	v1 := router.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	//对后续的路由使用中间件，用于用户鉴权
	v1.Use(middlewares.JWTAuthMiddleware())

	// 测试接口，在页面上显示当前登录的用户名
	v1.GET("/test", controller.TestHandle)
	// swagger接口文档
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return router
}
