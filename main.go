package main

import (
	"fmt"
	controller "xj/xapi-backend/controller"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	gconfig "xj/xapi-backend/g_config"
	glog "xj/xapi-backend/g_log"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/loadconfig"
	"xj/xapi-backend/middleware"
	"xj/xapi-backend/router"

	_ "xj/xapi-backend/rpc_api_service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func init() {
	// 实例化日志对象
	if logger, err := glog.SetupLogger(); err != nil {
		panic(err)
	} else {
		glog.Log = logger
	}

	gstore.TokenMemoryStore = make(map[string]bool)
	gstore.InterfaceFuncName = make(map[int64]string)

	loadconfig.LoadInterfaceFuncNameMap()
}

//	@title			xApi 项目
//	@version		1.0
//	@description	Api对外开放平台
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	小熊
//	@contact.url	https://github.com/xiaoxiongmao5
//	@contact.email	627516430@qq.com

//	@license.name	license.name
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8090

func main() {
	defer glog.Log.Writer().Close()

	// // 加载dubbo配置
	// if err := loadconfig.LoadDubboConfig(); err != nil {
	// 	glog.Log.Error("dubbo配置加载失败, err=:", err)
	// 	panic(err)
	// } else {
	// 	glog.Log.Info("dubbo配置加载成功")
	// }

	// 加载App配置数据
	if config, err := loadconfig.LoadAppConfig(); err != nil {
		glog.Log.Error("App配置加载失败, err=:", err)
		panic(err)
	} else {
		glog.Log.Info("App配置加载成功")
		gconfig.AppConfig = config
	}

	if dbcn, err := db.ConnectionPool(gconfig.AppConfig.Database.SavePath); err != nil {
		glog.Log.Error("数据库连接失败, err=", err)
		panic(err)
	} else {
		glog.Log.Infof("数据库连接成功,savePath=[%s]", gconfig.AppConfig.Database.SavePath)
		db.MyDB = dbcn
	}

	// 启动配置文件加载协程
	go loadconfig.LoadNewAppConfig()
	go loadconfig.RegisterServiceToNacos()

	r := gin.New()

	// 使用中间件格式化日志
	r.Use(middleware.CusLogger())

	// 使用中间件处理全局错误拦截
	r.Use(middleware.G_ErrorHandlerMiddleware())

	// 使用中间件来处理跨域请求，并允许携带 Cookie
	r.Use(middleware.CORSMiddleware())

	// 访问控制（黑名单）
	r.Use(middleware.FilterWithAccessControlInBlackIp())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/invoke", middleware.AuthMiddleware(), controller.InvokeInterface)

	router.UserRouter(r)
	router.InterfaceRouter(r)
	router.UserInterfaceInfoRouter(r)
	router.AnalysisRouter(r)

	port := fmt.Sprintf(":%d", gconfig.AppConfig.Server.Port)
	r.Run(port)
}
