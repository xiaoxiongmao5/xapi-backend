package router

import (
	"xj/xapi-backend/controller"
	"xj/xapi-backend/middleware"

	"github.com/gin-gonic/gin"
)

func ManagerRouter(r *gin.Engine) {
	router := r.Group("/manage", middleware.FilterWithAccessControlInAdminIp())
	router.GET("/config/ratelimit", controller.GetIPRateLimitConfig)
	router.PUT("/config/ratelimit", controller.UpdateIPRateLimitConfig)
}
