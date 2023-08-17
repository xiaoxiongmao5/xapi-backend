package controller

import (
	"fmt"
	"strconv"
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
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["ParameterError"], "参数错误"))
		return
	}
	_, err := service.CreateInterface(params)
	if err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["CreateUserFailed"], "接口创建失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "接口创建成功",
	})
}

//	@Summary		更新接口信息
//	@Description	更新接口信息
//	@Tags			接口相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.UpdateInterfaceParams	true	"接口信息"
//	@Success		200		{object}	object
//	@Router			/interface/update [post]
func UpdateInterface(c *gin.Context) {
	var params *models.UpdateInterfaceParams
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Printf("param UpdateInterfaceParams err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["ParameterError"], "参数错误"))
		return
	}
	err := service.UpdateInterface(params)
	if err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["CreateUserFailed"], "接口修改失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "接口修改成功",
	})
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
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["GetInterfaceListFailed"], "获取接口列表失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
		"data":   list,
	})
}

//	@Summary		删除接口
//	@Description	删除接口
//	@Tags			接口相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int		true	"接口Id"
//	@Success		200	{object}	object	"接口列表"
//	@Router			/interface/delete [get]
func DeleteInterface(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Printf("param DeleteInterface err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["ParameterError"], "参数错误"))
		return
	}
	ID := int64(id)
	err = service.DeleteInterface(ID)
	if err != nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["GetInterfaceListFailed"], "删除接口失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
	})
}
