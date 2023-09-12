package controller

import (
	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	glog "xj/xapi-backend/g_log"
	"xj/xapi-backend/models"
	"xj/xapi-backend/service"
	"xj/xapi-backend/utils"

	"github.com/gin-gonic/gin"
)

type GetUserInterfaceInfoByIdResponse struct {
	Result int                           `json:"result"`
	Msg    string                        `json:"msg"`
	Data   models.ValidUserInterfaceInfo `json:"data"`
}

//	@Summary		根据用户信息、接口id获取用户接口信息
//	@Description	根据用户信息、接口id获取用户接口信息
//	@Tags			用户与接口关系
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	path		int									true	"接口id"
//	@Success		200	{object}	GetUserInterfaceInfoByIdResponse	"接口列表"
//	@Router			/userinterface/{id} [get]
func GetUserInterfaceInfoById(c *gin.Context) {
	id := c.Param("id")
	if utils.AreEmptyStrings(id) {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	interfaceId, err := utils.String2Int64(id)
	if err != nil {
		glog.Log.Errorf("param id format err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	interfaceInfo, err := service.GetInterfaceInfoById(interfaceId)
	if err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口信息获取失败"))
		return
	}
	resData := models.Convert2ValidUserInterfaceInfo(interfaceInfo)
	userAccount, exists := c.Get("user_account")
	if !exists {
		ghandle.HandlerContextError(c, "user_account")
		return
	}
	// 拿到用户id
	userInfo, err := service.GetUserInfoByUserAccount(userAccount.(string))
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "用户不存在"))
		return
	}
	userInterfaceInfo, err := service.GetUserInterfaceInfoByUserIdAndInterfaceId(interfaceId, userInfo.ID)
	if err != nil {
		ghandle.HandlerSuccess(c, "接口信息获取成功", resData)
		return
	}
	resData.Totalnum = userInterfaceInfo.Totalnum
	resData.Leftnum = userInterfaceInfo.Leftnum
	ghandle.HandlerSuccess(c, "接口信息获取成功", resData)
}

//	@Summary		更新用户接口剩余调用次数
//	@Description	更新用户接口剩余调用次数
//	@Tags			用户与接口关系
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.UpdateInvokeLeftCountParams	true	"更新用户接口剩余调用次数"
//	@Success		200		{object}	object
//	@Router			/userinterface/update/leftcount [post]
func UpdateInvokeLeftCount(c *gin.Context) {
	var params *models.UpdateInvokeLeftCountParams
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("UpdateInvokeLeftCountParams err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 判断接口id是否存在
	if _, err := service.GetInterfaceInfoById(params.Interfaceinfoid); err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "接口不存在"))
		return
	}
	userAccount, exists := c.Get("user_account")
	if !exists {
		ghandle.HandlerContextError(c, "user_account")
		return
	}
	// 拿到用户id
	userInfo, err := service.GetUserInfoByUserAccount(userAccount.(string))
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "用户不存在"))
		return
	}

	// 更新用户接口剩余调用次数
	if _, err := service.UpdateInvokeLeftCount(params.Interfaceinfoid, userInfo.ID, params.Leftnum); err != nil {
		c.Error(gerror.NewAbortErr(int(enums.UpdateInvokeLeftCountFailed), err.Error()))
		return
	}

	ghandle.HandlerSuccess(c, "操作成功", nil)
}
