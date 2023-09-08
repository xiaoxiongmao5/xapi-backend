package router

import (
	"xj/xapi-backend/controller"
	"xj/xapi-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserInterfaceInfoRouter(r *gin.Engine) {
	router := r.Group("/userinterface", middleware.AuthMiddleware())

	router.GET("/:id", controller.GetUserInterfaceInfoById)
	router.POST("/update/leftcount", controller.UpdateInvokeLeftCount)
}
