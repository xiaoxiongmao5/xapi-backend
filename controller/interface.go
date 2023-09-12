package controller

import (
	"fmt"
	"reflect"
	"strconv"
	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	glog "xj/xapi-backend/g_log"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/models"
	"xj/xapi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxiongmao5/xapi-clientsdk/client"
)

//	@Summary		注册接口
//	@Description	注册接口
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.CreateInterfaceParams	true	"接口信息"
//	@Success		200		{object}	object
//	@Router			/interface/register [post]
func CreateInterface(c *gin.Context) {
	var params *models.CreateInterfaceParams
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("CreateInterfaceParams err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 注册接口
	if _, err := service.CreateInterface(params); err != nil {
		glog.Log.Errorf("service.CreateUser err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.CreateInterfaceFailed), "接口注册失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口注册成功", nil)
}

//	@Summary		更新接口信息
//	@Description	更新接口信息
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.UpdateInterfaceParams	true	"接口信息"
//	@Success		200		{object}	object
//	@Router			/interface/update [put]
func UpdateInterface(c *gin.Context) {
	var params *models.UpdateInterfaceParams
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("UpdateInterfaceParams err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}

	// 2. 更新接口信息
	if err := service.UpdateInterface(params); err != nil {
		glog.Log.Errorf("service.CreateUser err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.UpdateInterfaceFailed), "接口修改失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口修改成功", nil)
}

//	@Summary		获得所有接口列表
//	@Description	获取所有接口列表
//	@Tags			接口相关
//	@Produce		application/json
//	@Success		200	{object}	object	"接口列表"
//	@Router			/interface/list [get]
func ListInterface(c *gin.Context) {
	list, err := service.AllListInterfaces()
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表获取失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口列表获取成功", list)
}

//	@Summary		分页获得接口列表
//	@Description	分页获取接口列表
//	@Tags			接口相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			pageSize	query		int		true	"pageSize"
//	@Param			current		query		int		true	"current"
//	@Success		200			{object}	object	"接口列表"
//	@Router			/interface/pagelist [get]
func PageListInterface(c *gin.Context) {
	pageSize, err1 := strconv.Atoi(c.Query("pageSize"))
	current, err2 := strconv.Atoi(c.Query("current"))
	if err1 != nil || err2 != nil {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	list, err := service.PageListInterfaces(current, pageSize)
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表信息获取失败"))
		return
	}
	count, err := service.GetInterfaceListCount()
	if err != nil {
		c.Error(gerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表总数获取失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口列表获取成功", gin.H{
		"record": list,
		"total":  count,
	})
}

type GetInterfaceInfoByIdResponse struct {
	Result int                       `json:"result"`
	Msg    string                    `json:"msg"`
	Data   models.ValidInterfaceInfo `json:"data"`
}

//	@Summary		根据接口id获取接口信息
//	@Description	根据接口id获取接口信息
//	@Tags			接口相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	path		int								true	"接口id"
//	@Success		200	{object}	GetInterfaceInfoByIdResponse	"接口列表"
//	@Router			/interface/{id} [get]
func GetInterfaceInfoById(c *gin.Context) {
	if id := c.Param("id"); id == "" {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		glog.Log.Errorf("param id format err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	data, err := service.GetInterfaceInfoById(int64(id))
	if err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口信息获取失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口信息获取成功", data)
}

//	@Summary		删除接口
//	@Description	删除接口
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.IdRequest	true	"接口id"
//	@Success		200		{object}	object				"接口列表"
//	@Router			/interface/delete [delete]
func DeleteInterface(c *gin.Context) {
	var params *models.IdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("param IdRequest err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}
	// 2. 删除接口
	if err := service.DeleteInterface(params.ID); err != nil {
		c.Error(gerror.NewAbortErr(int(enums.DeleteInterfaceFailed), "接口删除失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口删除成功", nil)
}

//	@Summary		发布接口
//	@Description	发布接口
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.IdRequest	true	"接口id"
//	@Success		200		{object}	object
//	@Router			/interface/online [patch]
func OnlineInterface(c *gin.Context) {
	var params *models.IdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("param IdRequest err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}
	// 2. 检查接口是否可用（调用测试接口）

	// 3. 修改接口状态statuc=1
	if err := service.OnlineInterfaceStatus(params.ID); err != nil {
		glog.Log.Errorf("service.OnlineInterfaceStatus err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.OnlineInterfaceFailed), "接口发布失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口发布成功", nil)
}

//	@Summary		下线接口
//	@Description	下线接口
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.IdRequest	true	"接口id"
//	@Success		200		{object}	object
//	@Router			/interface/offline [patch]
func OfflineInterface(c *gin.Context) {
	var params *models.IdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("param IdRequest err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		glog.Log.Errorf("service.GetInterfaceInfoById err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}

	// 2. 修改接口状态statuc=0
	if err := service.OfflineInterfaceStatus(params.ID); err != nil {
		glog.Log.Errorf("service.OfflineInterfaceStatus err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.OfflineInterfaceFailed), "接口下线失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口下线成功", nil)
}

//	@Summary		调用接口
//	@Description	调用接口
//	@Tags			接口调用相关
//	@Accept			application/json
//	@Param			request	body	models.InvokeInterfaceParams	true	"调用接口参数"
//	@Router			/api/invoke [post]
func InvokeInterface(c *gin.Context) {
	var params *models.InvokeInterfaceParams
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("param InvokeInterfaceParams err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 默认传"{}"
	if params.Requestparams == "" {
		params.Requestparams = "{}"
	}

	// 获取用户的ak sk
	userAccount, exists := c.Get("user_account")
	if !exists {
		ghandle.HandlerContextError(c, "user_account")
		return
	}
	userInfo, err := service.GetUserInfoByUserAccount(userAccount.(string))

	// 检查接口剩余次数是否>0
	fullUserInterfaceInfo, err := service.GetFullUserInterfaceInfoByUserIdAndInterfaceId(params.ID, userInfo.ID)
	if err != nil {
		glog.Log.Errorf("service.GetFullUserInterfaceInfoByUserIdAndInterfaceId err=%v", err.Error())
		c.Error(gerror.NewAbortErr(int(enums.InterfaceNotExist), "接口剩余可调用次数不足"))
		return
	}
	// 检查接口剩余可调用次数
	if fullUserInterfaceInfo.Leftnum <= 0 {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "接口剩余可调用次数不足"))
		return
	}
	// 检查用户调用该接口是否被禁用
	if fullUserInterfaceInfo.BanStatus != enums.UserInterfaceStatusOk {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "该接口为禁用状态"))
		return
	}
	// 检查接口是否正常状态
	if fullUserInterfaceInfo.Status != enums.InterfaceStatusOnline {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "接口未上线"))
		return
	}

	// new一个客户端SDK
	clientsdk := client.NewClient(userInfo.Accesskey, userInfo.Secretkey)

	// 根据ID值从全局map中获取函数名
	funcName, res := gstore.InterfaceFuncName[params.ID]
	if res == false {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "接口暂未接入完成，敬请期待"))
		return
	}

	// 准备要传递的参数
	reflectArgs := make([]reflect.Value, 2)
	reflectArgs[0] = reflect.ValueOf(params.Requestparams)
	reflectArgs[1] = reflect.ValueOf(strconv.FormatInt(params.ID, 10))

	// 利用反射调用对应的函数
	method := reflect.ValueOf(clientsdk).MethodByName(funcName)
	if !method.IsValid() {
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "接口暂未接入完成，敬请期待"))
		return
	}
	result := method.Call(reflectArgs)
	// 如果没有返回值或提取值无效，返回错误
	if len(result) < 4 {
		c.Error(gerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回值格式校验失败"))
		return
	}
	flag := true
	// 提取 statusCode
	statusCode, ok := result[0].Interface().(int)
	if !ok {
		flag = false
	}
	// 提取 contentType
	contentType, ok := result[1].Interface().(string)
	if !ok {
		flag = false
	}
	// 提取 bodyBytes
	bodyBytes, ok := result[2].Interface().([]byte)
	if !ok {
		flag = false
	}
	// 提取 error
	bodyError, ok := result[3].Interface().(error)
	if ok {
		flag = false
	}

	glog.Log.WithFields(logrus.Fields{
		"res[0]statusCode":  fmt.Sprintf("%v", statusCode),
		"res[0]type":        fmt.Sprintf("%T", statusCode),
		"res[1]contentType": fmt.Sprintf("%v", contentType),
		"res[1]type":        fmt.Sprintf("%T", contentType),
		"res[2]bodyBytes":   fmt.Sprintf("%v", string(bodyBytes)),
		"res[2]type":        fmt.Sprintf("%T", bodyBytes),
		"res[3]bodyError":   fmt.Sprintf("%v", bodyError),
		"res[3]type":        fmt.Sprintf("%T", bodyError),
	}).Info("解析调用接口返回数据")

	if !flag {
		if bodyError != nil {
			// todo 这里可以降级处理
			c.Error(gerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回错误: "+bodyError.Error()))
		} else {
			c.Error(gerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回数据格式错误, 解析失败"))
		}
		return
	}

	// 使用提取的值调用 c.Data 将响应体内容直接返回给前端
	c.Data(statusCode, contentType, bodyBytes)
	return
}
