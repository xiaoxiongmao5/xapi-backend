package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "xj/xapi-backend/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
	"xj/xapi-backend/controller/user"
)

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

func setupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/login", controller.UserLogin)
    r.GET("/logout", controller.UserLogout)

	return r
}