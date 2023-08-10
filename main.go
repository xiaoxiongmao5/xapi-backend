package main

import (
	"fmt"
	controller "xj/xapi-backend/controller/user"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"
	"xj/xapi-backend/myerror"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	myerror.ResponseCodes = map[string]int{
		"Success":            0,
		"ParameterError":     1001,
		"AuthenticationFail": 1002,
		"UserNotExist":       2001,
		"UserExist":          2002,
		"CreateUserFailed":   2003,
		"UserPasswordError":  2004,
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

func setupRouter() *gin.Engine {
	r := gin.New()
	// 使用自定义的中间件
	r.Use(ErrorHandlerMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", controller.UserLogin)
	r.POST("/register", controller.UserRegister)
	r.GET("/logout", controller.UserLogout)

	return r
}
