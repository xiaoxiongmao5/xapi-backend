package controller

import (
	"fmt"
	"reflect"
	"strconv"
	"xj/xapi-backend/enums"
	ghandle "xj/xapi-backend/g_handle"
	"xj/xapi-backend/models"
	"xj/xapi-backend/myerror"
	"xj/xapi-backend/service"
	"xj/xapi-backend/store"

	"github.com/gin-gonic/gin"
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
		fmt.Printf("param CreateInterfaceParams err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 注册接口
	if _, err := service.CreateInterface(params); err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.CreateInterfaceFailed), "接口注册失败"))
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
		fmt.Printf("param UpdateInterfaceParams err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}

	// 2. 更新接口信息
	if err := service.UpdateInterface(params); err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.UpdateInterfaceFailed), "接口修改失败"))
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
		c.Error(myerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表获取失败"))
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
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	list, err := service.PageListInterfaces(current, pageSize)
	if err != nil {
		c.Error(myerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表信息获取失败"))
		return
	}
	count, err := service.GetInterfaceListCount()
	if err != nil {
		c.Error(myerror.NewAbortErr(int(enums.ListInterfaceFailed), "接口列表总数获取失败"))
		return
	}

	ghandle.HandlerSuccess(c, "接口列表获取成功", gin.H{
		"record": list,
		"total":  count,
	})
}

type ResponseWithData struct {
	Result int                           `json:"result"`
	Msg    string                        `json:"msg"`
	Data   models.ValidXapiInterfaceInfo `json:"data"`
}

//	@Summary		根据接口id获取接口信息
//	@Description	根据接口id获取接口信息
//	@Tags			接口相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	path		int					true	"接口id"
//	@Success		200	{object}	ResponseWithData	"接口列表"
//	@Router			/interface/{id} [get]
func GetInterfaceInfoById(c *gin.Context) {
	if id := c.Param("id"); id == "" {
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("param id format err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	data, err := service.GetInterfaceInfoById(int64(id))
	if err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口信息获取失败"))
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
		fmt.Printf("param IdRequest err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}
	// 2. 删除接口
	if err := service.DeleteInterface(params.ID); err != nil {
		c.Error(myerror.NewAbortErr(int(enums.DeleteInterfaceFailed), "接口删除失败"))
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
		fmt.Printf("param IdRequest err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}
	// 2. 检查接口是否可用（调用测试接口）

	// 3. 修改接口状态statuc=1
	if err := service.OnlineInterfaceStatus(params.ID); err != nil {
		fmt.Printf("service.OnlineInterfaceStatus err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.OnlineInterfaceFailed), "接口发布失败"))
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
		fmt.Printf("param IdRequest err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 1. 检查接口是否存在
	if _, err := service.GetInterfaceInfoById(params.ID); err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}

	// 2. 修改接口状态statuc=0
	if err := service.OfflineInterfaceStatus(params.ID); err != nil {
		fmt.Printf("service.OfflineInterfaceStatus err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.OfflineInterfaceFailed), "接口下线失败"))
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
		fmt.Printf("param InvokeInterfaceParams err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}
	// 检查接口ID是否小于等于0
	if params.ID <= 0 {
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}

	// 检查接口是否存在
	interfaceInfo, err := service.GetInterfaceInfoById(params.ID)
	if err != nil {
		fmt.Printf("service.GetInterfaceInfoById err=%v \n", err)
		c.Error(myerror.NewAbortErr(int(enums.InterfaceNotExist), "接口不存在"))
		return
	}

	// 检查接口是否正常状态
	if interfaceInfo.Status != int32(enums.InterfaceStatusOnline) {
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "接口未上线"))
		return
	}

	// 获取用户的ak sk
	userAccount, exists := c.Get("user_id")
	if !exists {
		ghandle.HandlerContextError(c, "user_id")
		return
	}
	userInfo, err := service.GetUserInfoByUserAccount(userAccount.(string))

	// new一个客户端SDK
	clientsdk := client.NewClient(userInfo.Accesskey, userInfo.Secretkey)

	// 根据ID值从全局map中获取函数名
	funcName, res := store.InterfaceFuncName[params.ID]
	if res == false {
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "接口暂未接入完成，敬请期待"))
		return
	}

	// 准备要传递的参数
	reflectArgs := make([]reflect.Value, 1)
	reflectArgs[0] = reflect.ValueOf(params.Requestparams)

	// 利用反射调用对应的函数
	method := reflect.ValueOf(clientsdk).MethodByName(funcName)
	if !method.IsValid() {
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "接口暂未接入完成，敬请期待"))
		return
	}
	result := method.Call(reflectArgs)
	// 如果没有返回值或提取值无效，返回错误
	if len(result) < 4 {
		c.Error(myerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回值格式校验失败"))
		return
	}
	// 提取 statusCode
	statusCode, ok := result[0].Interface().(int)
	if !ok {
		c.Error(myerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回无效的statusCode"))
		return
	}
	fmt.Printf("statusCode=%v \ttype=%T\n", statusCode, statusCode)

	// 提取 contentType
	contentType, ok := result[1].Interface().(string)
	if !ok {
		c.Error(myerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回无效的contentType"))
		return
	}
	fmt.Printf("contentType=%v \ttype=%T\n", contentType, contentType)

	// 提取 bodyBytes
	bodyBytes, ok := result[2].Interface().([]byte)
	if !ok {
		c.Error(myerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回无效的bodyBytes"))
		return
	}
	fmt.Printf("bodyBytes=%v \ttype=%T\n", string(bodyBytes), bodyBytes)

	// 提取 error
	bodyError, ok := result[3].Interface().(error)
	if ok {
		fmt.Printf("bodyError=%v  \ttype=%T\n", bodyError, bodyError)
		// todo 这里可以降级处理
		c.Error(myerror.NewAbortErr(int(enums.InvokeInterfaceFailed), "调用接口返回错误: "+bodyError.Error()))
		return
	}
	// 使用提取的值调用 c.Data 将响应体内容直接返回给前端
	c.Data(statusCode, contentType, bodyBytes)
	return
}
