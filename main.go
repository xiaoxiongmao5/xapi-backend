package main

import (
	"fmt"
	controller "xj/xapi-backend/controller"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	gconfig "xj/xapi-backend/g_config"
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
	// 加载dubbo配置
	if err := loadconfig.LoadDubboConfig(); err != nil {
		fmt.Println("LoadDubboConfig failed:", err)
		panic(err)
	}

	// 加载App配置数据
	if config, err := loadconfig.LoadAppConfig(); err != nil {
		fmt.Println("LoadAppConfig failed:", err)
		panic(err)
	} else {
		gconfig.AppConfig = config
	}

	// 从环境变量中获取 MySQL 连接信息
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// if !utils.AreEmptyStrings(dbHost, dbPort, dbUser, dbPassword, dbName) {
	// 	gconfig.AppConfig.Database.Host = dbHost
	// 	gconfig.AppConfig.Database.Port, _ = strconv.Atoi(dbPort)
	// 	gconfig.AppConfig.Database.Dbname = dbName
	// 	gconfig.AppConfig.Database.Username = dbUser
	// 	gconfig.AppConfig.Database.Password = dbPassword
	// } else {
	// 	fmt.Println("environment dbconfig error!!")
	// }

	// 构建 MySQL 连接字符串
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	gstore.TokenMemoryStore = make(map[string]bool)

	InitInterfaceFuncName()

	db.MyDB = db.ConnectionPool(gconfig.AppConfig.Database.SavePath)
}

func InitInterfaceFuncName() {
	gstore.InterfaceFuncName = make(map[int64]string)
	gstore.InterfaceFuncName = map[int64]string{
		1: "GetNameByGet",
		2: "GetNameByGet",
		3: "GetNameByPost",
	}
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

//	@host	localhost:8080

func main() {
	// 启动配置文件加载协程
	go loadconfig.LoadNewAppConfig()
	// go loadconfig.RegisterServiceToNacos()

	r := gin.New()
	// 使用自定义的中间件处理全局错误拦截
	r.Use(middleware.G_ErrorHandlerMiddleware())
	// 使用中间件来处理跨域请求，并允许携带 Cookie
	r.Use(middleware.CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/invoke", middleware.AuthMiddleware(), controller.InvokeInterface)

	router.UserRouter(r)
	router.InterfaceRouter(r)
	router.UserInterfaceInfoRouter(r)

	port := fmt.Sprintf(":%d", gconfig.AppConfig.Server.Port)
	r.Run(port)
}
