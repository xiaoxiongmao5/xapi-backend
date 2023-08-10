package main

import (
	"fmt"
	controller "xj/xapi-backend/controller/user"
	"xj/xapi-backend/db"
	_ "xj/xapi-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
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

func setupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", controller.UserLogin)
	r.POST("/register", controller.UserRegister)
	r.GET("/logout", controller.UserLogout)

	return r
}
