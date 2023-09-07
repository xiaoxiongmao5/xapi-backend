package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	gconfig "xj/xapi-backend/g_config"
	gstore "xj/xapi-backend/g_store"
	"xj/xapi-backend/models"
	"xj/xapi-backend/utils"

	"github.com/gin-gonic/gin"
)

// 根据Id 获取用户信息
func GetUserInfoById(id int64) (*dbsq.User, error) {
	// todo 这里是否需要new 新的链接地址
	q := dbsq.New(db.MyDB)
	// 创建一个 context.Context 对象
	ctx := context.Background()
	return q.GetUserInfoById(ctx, id)
}

// 根据userAccount 获得用户信息
func GetUserInfoByUserAccount(userAccount string) (*dbsq.User, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.GetUserInfoByUniUserAccount(ctx, userAccount)
}

// 根据accessKey 获得用户信息
func GetUserInfoByAccessKey(accessKey string) (*dbsq.User, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.GetUserInfoByUniAccessKey(ctx, accessKey)
}

// 创建账号
func CreateUser(param *models.UserRegisterParams) (sql.Result, error) {
	userAccount, userPassword, checkUserPassword := param.UserAccount, param.UserPassword, param.CheckUserPassword
	// 检验
	if utils.AreEmptyStrings(userAccount, userPassword, checkUserPassword) {
		return nil, errors.New("参数为空")
	}
	if length := len(userAccount); length < 4 || length > 16 {
		return nil, errors.New("用户账号长度不符合规定,长度要求4~16位")
	}
	if length := len(userPassword); length < 6 || length > 16 {
		return nil, errors.New("用户密码长度不符合规定,长度要求6~16位")
	}
	// 密码和校验密码相同
	if !utils.CheckSame[string]("校验两次输入的密码一致", userPassword, checkUserPassword) {
		return nil, errors.New("两次输入的密码不一致")
	}
	// 账号不能重复
	if _, err := GetUserInfoByUserAccount(userAccount); err == nil {
		return nil, errors.New("账户已存在")
	}
	// 将密码进行哈希
	hashPassword, err := utils.HashPasswordByBcrypt(userPassword)
	if err != nil {
		fmt.Printf("utils.HashByBcrypt err=%v \n", err)
		return nil, err
	}
	// 分配accessKey,secretKey
	accessKey := utils.HashBySHA256WithSalt(userAccount+utils.GenetateRandomString(5), gconfig.SALT)
	secretKey := utils.HashBySHA256WithSalt(userAccount+utils.GenetateRandomString(8), gconfig.SALT)

	params := &dbsq.CreateUserParams{
		Useraccount:  userAccount,
		Userpassword: hashPassword,
		Accesskey:    accessKey,
		Secretkey:    secretKey,
	}
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.CreateUser(ctx, params)
}

// 删除token
func DeleteToken(c *gin.Context) {
	// 从cookie拿到token
	tokenCookie, err := c.Cookie("token")
	if err != nil || tokenCookie == "" {
		return
	}

	// 从服务端删除该token
	delete(gstore.TokenMemoryStore, tokenCookie)
}
