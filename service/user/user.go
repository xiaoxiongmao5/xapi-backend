package service_user

import (
	"context"
	"database/sql"
	"fmt"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/utils"
)

// 获得用户信息
func GetUserInfo(userAccount string) (*dbsq.User, error) {
	// todo 这里是否需要new 新的链接地址
	q := dbsq.New(db.MyDB)
	// 创建一个 context.Context 对象
	ctx := context.Background()
	return q.GetUserInfo(ctx, userAccount)
}

// 创建账号
func CreateUser(params *dbsq.CreateUserParams) (sql.Result, error) {
	// 将密码进行哈希
	HashPassword, err := utils.HashPassword(params.Userpassword)
	if err != nil {
		fmt.Printf("utils.HashPassword err=%v \n", err)
		return nil, err
	}
	params.Userpassword = HashPassword
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.CreateUser(ctx, params)
}
