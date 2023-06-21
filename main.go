package main

import (
	"context"
	"flag"
	"fmt"
	"gin-web-scaffolding/dao/mysql"
	"gin-web-scaffolding/logger"
	"gin-web-scaffolding/routes"
	"gin-web-scaffolding/settings"
	"gin-web-scaffolding/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @file      : main.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// @title GinWeb脚手架项目
// @version v1.0.0
// @description GinWeb脚手架项目
// @contact.name 八宝糖
// @contact.email 1013269096@qq.com
// @host 127.0.0.1
// @BasePath /api/v1/
func main() {
	// 通过命令行参数指定配置文件的路经
	var configFilePath string
	var initDataBase bool
	flag.StringVar(&configFilePath, "config", "./config/config.yaml", "请指定配置文件路经")
	flag.BoolVar(&initDataBase, "initdb", false, "是否初始化DB")
	flag.Parse() // 解析命令行参数

	// 从配置文件中加载配置信息
	if err := settings.Init(configFilePath); err != nil {
		fmt.Printf("init settings failed, err:%#v\n", err)
		return
	}

	//  初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%#v\n", err)
		return
	}

	// 代码运行到这，就可以将信息记录到日志中了，因此在这里注册一个延迟调用
	defer func() {
		if err := zap.L().Sync(); err != nil {
			fmt.Printf("zap sync failed, err%#v\n", err)
		}
	}()

	// 初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%#v", err)
		return
	}

	// 如果initDataBase=true，则进行初始化数据库操作，初始化完成后程序退出
	if initDataBase {
		// 初始化数据库
		if err := mysql.Migrate(); err != nil {
			fmt.Printf("init database failed, er:%#v\n", err)
		}
		return
	}

	//  初始化雪花算法
	if err := utils.SnowFlakeInit(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%#v\n", err)
		return
	}

	// 初始化gin框架内置校验器使用的翻译器
	if err := utils.InitValidatorTrans("zh"); err != nil {
		fmt.Printf("init vaildtor trans failed, err:%#v\n", err)
		return
	}

	// 注册路由
	router := routes.SetupRoutes(settings.Conf.Mode, settings.Conf.AllowOrigins, settings.Conf.AllowMethods)

	// 启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen failed,", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅的关闭服务器，为关闭服务器操作设置一个5s的超时
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //此处不会阻塞
	<-quit                                               //阻塞在此，只有当接收到了上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server...")

	// 创建一个5s的超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5s内优雅的关闭服务(将未处理的请求处理完后再关闭服务)，如果超过了5s就超时退出强制关闭服务器了
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting...")

}
