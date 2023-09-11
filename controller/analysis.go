package controller

import (
	"fmt"
	"strconv"
	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	"xj/xapi-backend/models"
	"xj/xapi-backend/service"

	"github.com/gin-gonic/gin"
)

type ListTopNOfInterfaceInvokeCountResponse struct {
	Result int                                         `json:"result"`
	Msg    string                                      `json:"msg"`
	Data   []models.ValidTopNOfInterfaceInvokeCountRow `json:"data"`
}

//	@Summary		获取接口调用次数TopN的信息列表
//	@Description	通过用户接口信息表查询调用次数最多的接口ID，再关联查询接口详细信息
//	@Tags			分析统计用户与接口关系
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			n	query		int										true	"TOP N"
//	@Success		200	{object}	ListTopNOfInterfaceInvokeCountResponse	"接口列表"
//	@Router			/analysis/top/interface/invoke [get]
func ListTopNOfInterfaceInvokeCount(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	if limit <= 0 {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误, n不能小于0"))
		return
	}
	data, err := service.ListTopNOfInterfaceInvokeCount(int32(limit))
	if err != nil {
		fmt.Printf("service.ListTopNOfInterfaceInvokeCount err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口调用次数TopN的信息列表获取失败"))
		return
	}
	resData := service.ConvertSliceToValidTopNOfInterfaceInvokeCountRow(data)

	ghandle.HandlerSuccess(c, "接口调用次数TopN的信息列表获取成功", resData)
}
