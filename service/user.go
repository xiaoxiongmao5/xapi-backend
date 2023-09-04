package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"xj/xapi-backend/config"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/models"
	"xj/xapi-backend/store"
	"xj/xapi-backend/utils"

	"github.com/gin-gonic/gin"
)

// 根据userAccount 获得用户信息
func GetUserInfoByUserAccount(userAccount string) (*dbsq.User, error) {
	// todo 这里是否需要new 新的链接地址
	q := dbsq.New(db.MyDB)
	// 创建一个 context.Context 对象
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
func CreateUser(param *models.CreateUserParamsJSON) (sql.Result, error) {
	userAccount, userPassword, checkUserPassword := param.UserAccount, param.UserPassword, param.CheckUserPassword
	// 检验
	if utils.AreEmptyStrings(userAccount, userPassword, checkUserPassword) {
		return nil, errors.New("参数为空")
	}
	if len(userAccount) < 4 {
		return nil, errors.New("用户账号过短")
	}
	if len(userPassword) < 8 || len(checkUserPassword) < 8 {
		return nil, errors.New("用户密码过短")
	}
	// 密码和校验密码相同
	if userPassword != checkUserPassword {
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
	var rand5, rand8 string

	if rand5, err = utils.GenerateRandomKey(5); err != nil {
		fmt.Printf("utils.GenerateRandomKey err=%v \n", err)
		return nil, err
	}
	if rand8, err = utils.GenerateRandomKey(8); err != nil {
		fmt.Printf("utils.GenerateRandomKey err=%v \n", err)
		return nil, err
	}
	accessKey := utils.HashBySHA256WithSalt(userAccount+rand5, config.SALT)
	secretKey := utils.HashBySHA256WithSalt(userAccount+rand8, config.SALT)

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
	delete(store.TokenMemoryStore, tokenCookie)
}
