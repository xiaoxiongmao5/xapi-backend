package controller

import (
	"xj/xapi-backend/myerror"
	"xj/xapi-backend/service"

	"github.com/gin-gonic/gin"
)

func CreateInterface() {
	// dbsq.CreateInterface()
}

//	@Summary		获得接口列表
//	@Description	获取接口列表信息
//	@Tags			接口相关
//	@Produce		application/json
//	@Success		200	{object}	object	"接口列表"
//	@Router			/interface/list [get]
func ListInterface(c *gin.Context) {
	list, err := service.ListInterfaces()
	if err != nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["GetInterfaceListFailed"], "接口列表获取失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
		"data":   list,
	})
}
