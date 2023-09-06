package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"xj/xapi-backend/enums"
	gerror "xj/xapi-backend/g_error"
	ghandle "xj/xapi-backend/g_handle"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/models"
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
		c.Error(gerror.NewAbortErr(int(enums.ParameterError), "参数错误"))
		return
	}

	// 注册用户
	if _, err := service.CreateUser(params); err != nil {
		fmt.Printf("service.CreateUser err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.CreateUserFailed), err.Error()))
		return
	}

	ghandle.HandlerSuccess(c, "账号创建成功", nil)
}

//	@Summary		获得用户信息
//	@Description	根据用户的 Cookie 获取用户信息
//	@Tags			用户相关
//	@Produce		application/json
//	@Success		200	{object}	object	"用户信息"
//	@Router			/user/uinfo [get]
func GetUserInfoByUserAccount(c *gin.Context) {
	useraccount, exists := c.Get("user_id")
	if !exists {
		ghandle.HandlerContextError(c, "user_id")
		return
	}

	userInfo, err := service.GetUserInfoByUserAccount(useraccount.(string))
	if err != nil {
		fmt.Printf("q.GetUserInfoByUserAccount err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.UserNotExist), "用户不存在"))
		return
	}

	ghandle.HandlerSuccess(c, "success", models.ConvertToNormalUser(userInfo))
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
	// // 检查用户是否已经登录
	// tokenCookie, err := c.Cookie("token")
	// if err == nil && tokenCookie != "" {
	// 	// 已经登录，直接返回登录成功
	// ghandle.HandlerSuccess(c, "Already logged in", nil)
	// 	return
	// }
	service.DeleteToken(c)

	useraccount := c.PostForm("useraccount")
	userpassword := c.PostForm("userpassword")

	// 用户是否存在（获取用户信息）
	userInfo, err := service.GetUserInfoByUserAccount(useraccount)
	if err != nil {
		fmt.Printf("q.GetUserInfoByUserAccount err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.UserNotExist), "账号不存在"))
		return
	}
	fmt.Printf("拿到用户信息了%v \n", userInfo)

	// 验证密码是否正常
	if err := utils.CheckHashPasswordByBcrypt(userInfo.Userpassword, userpassword); err != nil {
		fmt.Printf("CheckHashPasswordByBcrypt err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.UserPasswordError), "账号不存在或者密码验证错误"))
		return
	}

	// 生成token
	token, err := utils.GenerateToken(useraccount, userInfo.Userrole)
	if err != nil {
		fmt.Printf("utils.GenerateToken err=%v \n", err)
		c.Error(gerror.NewAbortErr(int(enums.GenerateTokenFailed), err.Error()))
		return
	}

	// 存储token
	gstore.TokenMemoryStore[token] = true

	// 返回token到前端
	domain, _ := utils.GetDomainFromReferer(c.Request.Referer())
	c.SetCookie("token", token, 3600, "/", domain, false, true)

	ghandle.HandlerSuccess(c, "success", models.ConvertToNormalUser(userInfo))
}

//	@Summary		用户退出
//	@Description	用户退出
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/user/logout [get]
func UserLogout(c *gin.Context) {
	// 服务端删除token
	service.DeleteToken(c)

	// // 从前端删除该token的cookie
	// domain, _ := utils.GetDomainFromReferer(c.Request.Referer())
	// c.SetCookie("token", "", -1, "/", domain, false, true)

	ghandle.HandlerSuccess(c, "success", nil)
}
