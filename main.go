package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"xj/xapi-backend/config"
	controller "xj/xapi-backend/controller"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	"xj/xapi-backend/enums"
	"xj/xapi-backend/myerror"
	"xj/xapi-backend/store"
	"xj/xapi-backend/utils"

	_ "xj/xapi-backend/rpc_api_service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func init() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)

	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := dubboConfig.Load(); err != nil {
		panic(err)
	}
	store.TokenMemoryStore = make(map[string]bool)
	InitInterfaceFuncName()
	db.MyDB = db.ConnectionPool("root:@/xapi?charset=utf8&parseTime=true")
}

func InitInterfaceFuncName() {
	store.InterfaceFuncName = make(map[int64]string)
	store.InterfaceFuncName = map[int64]string{
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
	fmt.Println("hi xj")
	r := setupRouter()
	r.Run(":8090")

}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 判断上层业务抛出的错误类型
		if err := c.Errors.Last(); err != nil {
			if abortError, ok := err.Err.(*myerror.AbortError); ok {
				// 生成错误响应并终止请求处理
				c.JSON(http.StatusOK, gin.H{
					"result": abortError.Code,
					"msg":    abortError.Message,
				})
				c.Abort()
				return
			}
		}
	}
}

// CORSMiddleware 是处理跨域请求的中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域请求的来源域，这里需要设置为请求的 Origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))

		// 允许的 HTTP 方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// 允许的请求标头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Nonce, timestamp, accessKey, sign")

		// 允许携带 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求（OPTIONS 请求），直接返回成功状态
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取当前的 Token
		tokenCookie, err := c.Cookie("token")
		if err != nil || tokenCookie == "" {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Error(myerror.NewAbortErr(int(enums.Unauthorized), "Unauthorized"))
			c.Abort()
			return
		}

		// 验证当前 Token
		token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		})
		if err != nil || !token.Valid {
			c.Error(myerror.NewAbortErr(int(enums.Unauthorized), "Unauthorized"))
			c.Abort()
			return
		}

		// 从 Token 中获取用户信息
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Error(myerror.NewAbortErr(int(enums.Unauthorized), "Unauthorized"))
			c.Abort()
			return
		}

		// 重新生成 Token，并更新有效期
		userID := claims["user_id"].(string)
		userRole := claims["user_role"].(string)
		newToken, err := utils.GenerateToken(userID, userRole)
		if err != nil {
			c.Error(myerror.NewAbortErr(int(enums.GenerateTokenFailed), err.Error()))
			c.Abort()
			return
		}

		// 删除旧的 token
		delete(store.TokenMemoryStore, tokenCookie)

		// 更新内存中的 token 数据
		store.TokenMemoryStore[newToken] = true

		// 将新的 token 返回给前端
		domain, _ := utils.GetDomainFromReferer(c.Request.Referer())
		c.SetCookie("token", newToken, 3600, "/", domain, false, true)

		// 在此可以将 claims 中的用户信息保存到上下文中，供后续处理使用
		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["user_role"])

		c.Next()
	}
}

func AdminAuthMiddleware(c *gin.Context) {
	// 从上下文中获取用户信息
	userrole, exists := c.Get("user_role")
	if !exists || userrole.(string) != "admin" {
		c.Error(myerror.NewAbortErr(int(enums.NotAdminRole), "无权限"))
		c.Abort()
		return
	}
	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.New()
	// 使用自定义的中间件处理全局错误拦截
	r.Use(ErrorHandlerMiddleware())
	// 使用中间件来处理跨域请求，并允许携带 Cookie
	r.Use(CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/user")
	userRouter.POST("/login", controller.UserLogin)
	userRouter.POST("/register", controller.UserRegister)
	userRouter.GET("/logout", AuthMiddleware(), controller.UserLogout)
	userRouter.GET("/uinfo", AuthMiddleware(), controller.GetUserInfoByUserAccount)

	interfaceRouter := r.Group("/interface", AuthMiddleware())
	interfaceRouter.GET("/:id", controller.GetInterfaceInfoById)
	interfaceRouter.GET("/list", controller.ListInterface)
	interfaceRouter.GET("/pagelist", controller.PageListInterface)
	interfaceRouter.POST("/register", controller.CreateInterface)
	interfaceRouter.PUT("/update", controller.UpdateInterface)
	interfaceRouter.DELETE("/delete", controller.DeleteInterface)

	interfaceRouter.PATCH("/online", AdminAuthMiddleware, controller.OnlineInterface)
	interfaceRouter.PATCH("/offline", AdminAuthMiddleware, controller.OfflineInterface)

	r.POST("/api/invoke", AuthMiddleware(), controller.InvokeInterface)

	return r
}
