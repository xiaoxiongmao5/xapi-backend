package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"xj/xapi-backend/models"
	"xj/xapi-backend/myerror"
	service_user "xj/xapi-backend/service/user"
	"xj/xapi-backend/utils"

	_ "github.com/go-sql-driver/mysql" //导入
	// "log"
)

// @Summary		用户注册
// @Description	用户注册
// @Tags			用户相关
// @Accept			application/json
// @Produce		json
// @Param			username		body		string	false	"用户昵称"
// @Param			useraccount		body		string	true	"账号"
// @Param			useravatar		body		string	false	"用户头像"
// @Param			gender			body		int		false	"性别"
// @Param			userpassword	body		string	true	"密码"
// @Param			accesskey		body		string	true	"accessKey"
// @Param			secretkey		body		string	true	"secretKey"
// @Success		200				{object}	string
// @Router			/register [post]
func UserRegister(c *gin.Context) {
	var params *models.CreateUserParamsJSON
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Printf("q.CreateUser err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["ParameterError"], "参数错误"))
		return
	}
	// 获取该账号是否存在过
	_, err := service_user.GetUserInfo(params.Useraccount)
	if err == nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserExist"], "该账号名已被使用，请输入新的账号名"))
		return
	}
	params.Userrole = "user"
	_, err = service_user.CreateUser(models.ConvertToCreateUserParamsJSON(params))
	// fmt.Println("res=", res)
	if err != nil {
		fmt.Printf("service_user.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["CreateUserFailed"], "账号创建失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "账号创建成功",
	})
}

// @Summary		用户登录
// @Description	用户登录
// @Tags			用户相关
// @Accept			application/x-www-form-urlencoded
// @Produce		json
// @Param			useraccount		formData	string	true	"账户名称"
// @Param			userpassword	formData	string	true	"密码"
// @Success		200				{object}	string
// @Router			/login [post]
func UserLogin(c *gin.Context) {
	useraccount := c.PostForm("useraccount")
	userpassword := c.PostForm("userpassword")

	userInfo, err := service_user.GetUserInfo(useraccount)
	if err != nil {
		fmt.Printf("q.GetUserInfo err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserNotExist"], "用户不存在"))
		return
	}
	fmt.Printf("拿到用户信息了%v \n", userInfo)
	// 密码验证
	err = utils.ComparePassword(userInfo.Userpassword, userpassword)
	if err != nil {
		fmt.Printf("HashPassword err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserPasswordError"], "账号不存在或者密码验证错误"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
		"data":   models.ConvertToNormalUser(userInfo),
	})
}

// @Summary		用户退出
// @Description	用户退出
// @Tags			用户相关
// @Accept			application/json
// @Produce		application/json
// @Param			useraccount	query		string	true	"账户名称"
// @Success		200			{object}	string
// @Router			/logout [get]
func UserLogout(c *gin.Context) {
	useraccount := strings.Trim(c.Query("useraccount"), " ")
	fmt.Println("useraccount=", useraccount)
	// 去掉客户端的cookie
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
	})
}
