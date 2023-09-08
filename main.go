package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	controller "xj/xapi-backend/controller"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	gconfig "xj/xapi-backend/g_config"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/middleware"
	"xj/xapi-backend/router"

	_ "xj/xapi-backend/rpc_api_service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func init() {
	// 使用命令行参数来指定配置文件路径
	dubbogoConfigFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *dubbogoConfigFile)

	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := dubboConfig.Load(); err != nil {
		panic(err)
	}

	// 打开项目配置文件
	appConfigFile, err := os.Open("conf/appconfig.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		panic(err)
	}
	defer appConfigFile.Close()

	// 解码配置文件内容到结构体
	decoder := json.NewDecoder(appConfigFile)
	if err = decoder.Decode(&gconfig.AppConfig); err != nil {
		fmt.Println("Error decoding config file:", err)
		panic(err)
	}

	gstore.TokenMemoryStore = make(map[string]bool)
	InitInterfaceFuncName()
	dbcfg := gconfig.AppConfig.Database
	savePath := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=true", dbcfg.Username, dbcfg.Dbname)
	db.MyDB = db.ConnectionPool(savePath)
	// db.MyDB = db.ConnectionPool("root:@/xapi?charset=utf8&parseTime=true")
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
	// r := setupRouter()
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
