package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "errors"
	// "xj/xapi-backend/dbsq"
)

type ResCode map[int]string

 
// _ResponsePostList 
type _ResponsePostList[T any] struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    T `json:"data"`    // 数据
}

type User struct {
	// 账号
	Useraccount string
	// 密码
	Userpassword string
}
//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			用户相关
//	@Accept			json
//	@Produce		json
//	@Param			Useraccount		formData	string	true	"账户名"
//	@Param			Userpassword	formData	string	true	"密码"
//	@Success		200				{object}	string
//	@Router			/login [post]
func UserLogin(c *gin.Context) {
	var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"result": 1, 
			"msg": "fail",
		})
        return
    }
	c.JSON(200, gin.H{
		"result": 0, 
		"msg": "success",
		"data"	: "用户登录 ok",
	})
}

//	@Summary		用户退出
//	@Description	用户退出
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Useraccount	query		string	true	"账户名"
//	@Success		200			{object}	_ResponsePostList[string]
//	@Router			/logout [get]
func UserLogout(c *gin.Context) {
    user := c.Query("Useraccount")
	// if user == "" {
	// 	fmt.Println(errors.New("用户不存在"))
	// 	c.JSON(200, gin.H{
	// 		"result": 1, 
	// 		"msg": "fail",
	// 	})
    //     return
	// }
	c.JSON(200, gin.H{
		"result": 0, 
		"msg": "success",
		"data"	: fmt.Sprintf("用户%s退出 ok", user),
	})
}
