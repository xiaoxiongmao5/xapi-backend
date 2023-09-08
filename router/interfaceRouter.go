package router

import (
	"xj/xapi-backend/controller"
	"xj/xapi-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InterfaceRouter(r *gin.Engine) {
	router := r.Group("/interface", middleware.AuthMiddleware())

	router.GET("/:id", controller.GetInterfaceInfoById)
	router.GET("/list", controller.ListInterface)
	router.GET("/pagelist", controller.PageListInterface)
	router.POST("/register", controller.CreateInterface)
	router.PUT("/update", controller.UpdateInterface)
	router.DELETE("/delete", controller.DeleteInterface)

	router.PATCH("/online", middleware.AdminMiddleware(), controller.OnlineInterface)
	router.PATCH("/offline", middleware.AdminMiddleware(), controller.OfflineInterface)
}
