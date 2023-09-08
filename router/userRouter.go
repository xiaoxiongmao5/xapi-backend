package router

import (
	"xj/xapi-backend/controller"
	"xj/xapi-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	router := r.Group("/user")

	router.POST("/login", controller.UserLogin)
	router.POST("/register", controller.UserRegister)
	router.GET("/logout", middleware.AuthMiddleware(), controller.UserLogout)
	router.GET("/uinfo", middleware.AuthMiddleware(), controller.GetUserInfoByUserAccount)
}
