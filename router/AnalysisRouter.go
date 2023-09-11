package router

import (
	"xj/xapi-backend/controller"

	"github.com/gin-gonic/gin"
)

func AnalysisRouter(r *gin.Engine) {
	router := r.Group("/analysis")

	router.GET("/top/interface/invoke", controller.ListTopNOfInterfaceInvokeCount)
}
