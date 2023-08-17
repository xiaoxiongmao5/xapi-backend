package main

import (
	"fmt"
	controller "xj/xapi-backend/controller"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	"xj/xapi-backend/myerror"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	myerror.ResponseCodes = map[string]int{
		"Success":                0,
		"ParameterError":         1001,
		"AuthenticationFail":     1002,
		"UserNotExist":           2001,
		"UserExist":              2002,
		"CreateUserFailed":       2003,
		"UserPasswordError":      2004,
		"GetInterfaceListFailed": 3001,
	}
	db.MyDB = db.ConnectionPool("root:@/xapi?charset=utf8&parseTime=true")

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
	r.Run(":8080")

}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 判断上层业务抛出的错误类型
		if err := c.Errors.Last(); err != nil {
			if abortError, ok := err.Err.(*myerror.AbortError); ok {
				// 生成错误响应并终止请求处理
				c.JSON(200, gin.H{
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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 允许的请求标头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 允许携带 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求（OPTIONS 请求），直接返回成功状态
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
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
	userRouter.GET("/logout", controller.UserLogout)
	userRouter.GET("/uinfo", controller.GetUserInfo)

	interfaceRouter := r.Group("/interface")
	interfaceRouter.GET("/list", controller.ListInterface)
	interfaceRouter.POST("/register", controller.CreateInterface)
	interfaceRouter.POST("/update", controller.UpdateInterface)
	interfaceRouter.GET("/delete", controller.DeleteInterface)

	return r
}
