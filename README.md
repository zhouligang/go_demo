# GinWeb脚手架

#### 介绍
基于go语言的gin web框架实现的web开发的脚手架。

###### 使用到的主要技术点
1. 认证使用jwt，双token模式(accessToken、refreshToken)
2. 数据库操作使用gorm
3. 限流使用令牌桶
4. 用户id生成使用的是雪花算法
5. 日志管理使用zap
6. 配置文件管理使用viper
7. 代码热监听使用air
8. 接口文档生成使用swagger

#### 软件架构
软件各模块代码说明
1. main.go 程序的主入口
2. config 存放程序的配置文件信息
3. settings 读取配置文件中的信息
4. routes 注册的路由信息
5. middlewares 存放自定义中间件代码
6. utils 存放一些公共方法
7. models 存放数据库模型结构体、返回或请求参数的结构体等
8. logger 自定义的日志处理方法
9. controller 控制层代码，客户端请求后的处理函数
10. logic 逻辑层代码，处理具体的业务逻辑
11. dao 对数据库进行处理的代码
12. sqls 存放一些sql语句代码

#### 安装教程

go代码热监听
```shell 
1. go get -u github.com/cosmtrek/air
2. 编译：go build
3. 将编译后的ari拷贝到GOPATH/bin下
```

生成api文档
```shell 
1. go get -u github.com/swaggo/gin-swagger
2. go get -u github.com/swaggo/swag
3. 编译：进入cmd目录，执行go build
4. 将编译后的swag拷贝到GOPATH/bin下
5. go get -u github.com/swaggo/files
6. swag init
```

#### 使用说明
```shell 
1. go mod tidy
2. go run main.go
```
3. 使用postman对结果进行测试

#### 参与贡献
<font color="#0000dd">对于代码中的错误，还请批评指</font><br />
1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
