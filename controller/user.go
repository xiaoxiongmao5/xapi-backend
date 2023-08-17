package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"xj/xapi-backend/models"
	"xj/xapi-backend/myerror"
	service "xj/xapi-backend/service"
	"xj/xapi-backend/utils"

	_ "github.com/go-sql-driver/mysql" //导入
	// "log"
)

//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.CreateUserParamsJSON	true	"注册信息"
//	@Success		200		{object}	object
//	@Router			/user/register [post]
func UserRegister(c *gin.Context) {
	var params *models.CreateUserParamsJSON
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Printf("param CreateUserParamsJSON err=%v \n", err.Error())
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["ParameterError"], "参数错误"))
		return
	}
	// 获取该账号是否存在过
	_, err := service.GetUserInfo(params.Useraccount)
	if err == nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserExist"], "该账号名已被使用，请输入新的账号名"))
		return
	}
	params.Userrole = "user"
	_, err = service.CreateUser(models.ConvertToCreateUserParamsJSON(params))
	// fmt.Println("res=", res)
	if err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["CreateUserFailed"], "账号创建失败"))
		return
	}
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "账号创建成功",
	})
}

//	@Summary		获得用户信息
//	@Description	根据用户的 Cookie 获取用户信息
//	@Tags			用户相关
//	@Produce		application/json
//	@Success		200	{object}	object	"用户信息"
//	@Router			/user/uinfo [get]
func GetUserInfo(c *gin.Context) {
	// 从请求中获取 Cookie
	cookie, err := c.Cookie("token")
	if err != nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserNotExist"], "用户信息获取失败"))
		return
	}
	// // 这里假设有效的 token 是 "example_token"
	// if cookie != "example_token" {
	// 	c.String(http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }
	userInfo, err := service.GetUserInfo(cookie)
	if err != nil {
		fmt.Printf("q.GetUserInfo err=%v \n", err)
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserNotExist"], "用户不存在"))
		return
	}
	fmt.Printf("拿到用户信息了%v \n", userInfo)
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
		"data":   models.ConvertToNormalUser(userInfo),
	})
}

//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			用户相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			request	formData	models.UserLoginParamsJSON	true	"账号密码"
//	@Success		200		{object}	object
//	@Router			/user/login [post]
func UserLogin(c *gin.Context) {
	useraccount := c.PostForm("useraccount")
	userpassword := c.PostForm("userpassword")

	userInfo, err := service.GetUserInfo(useraccount)
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
	token := useraccount
	// 设置 Cookie
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
		"data":   models.ConvertToNormalUser(userInfo),
	})
}

//	@Summary		用户退出
//	@Description	用户退出
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/user/logout [get]
func UserLogout(c *gin.Context) {
	// 从请求中获取 Cookie
	cookie, err := c.Cookie("token")
	if err != nil {
		c.Error(myerror.NewAbortErr(myerror.ResponseCodes["UserNotExist"], "用户信息获取失败"))
		return
	}
	// // 这里假设有效的 token 是 "example_token"
	// if cookie != "example_token" {
	// 	c.String(http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	fmt.Println("token=", cookie)
	// 去掉客户端的cookie
	c.JSON(200, gin.H{
		"result": 0,
		"msg":    "success",
	})
}
