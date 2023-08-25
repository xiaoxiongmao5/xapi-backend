package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"xj/xapi-backend/enums"
	"xj/xapi-backend/models"
	"xj/xapi-backend/myerror"
	"xj/xapi-backend/service"

	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口注册成功",
	})
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口修改成功",
	})
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口列表获取成功",
		"data":   list,
	})
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口列表获取成功",
		"data": gin.H{
			"record": list,
			"total":  count,
		},
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口信息获取成功",
		"data":   data,
	})
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
	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口删除成功",
	})
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

	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口发布成功",
	})
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

	c.JSON(http.StatusOK, gin.H{
		"result": 0,
		"msg":    "接口下线成功",
	})
}

//	@Summary		调用接口
//	@Description	调用接口
//	@Tags			接口相关
//	@Accept			application/json
//	@Param			request	body	models.InvokeInterfaceParams	true	"调用接口参数"
//	@Router			/interface/invoke [post]
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
		c.Error(myerror.NewAbortErr(int(enums.ParameterError), "用户不存在"))
		return
	}
	userInfo, err := service.GetUserInfo(userAccount.(string))
	accesskey := userInfo.Accesskey
	secretkey := userInfo.Secretkey
	fmt.Println("accesskey=", accesskey)
	fmt.Println("secretkey=", secretkey)

	// 构建 URL 参数
	forwardHandlerParams := url.Values{}
	forwardHandlerParams.Add("name", params.Requestparams)
	forwardHandler(c, forwardHandlerParams)
}

func forwardHandler(c *gin.Context, params url.Values) {
	requestURL := "http://localhost:8002/check" + "?" + params.Encode()
	// 发起请求到第三方接口
	response, err := http.Get(requestURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// 将响应体内容直接返回给前端
	c.Data(response.StatusCode, response.Header.Get("Content-Type"), bodyBytes)
}
