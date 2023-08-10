package models

import (
	"database/sql"
	"time"
	"xj/xapi-backend/dbsq"
)

// 登录获取用户信息显示
type ShowUserJSON struct {
	// id
	ID int64 `json:"id"`
	// 用户昵称
	Username string `json:"username"`
	// 账号
	Useraccount string `json:"useraccount"`
	// 用户头像
	Useravatar string `json:"useravatar"`
	// 性别
	Gender int32 `json:"gender"`
	// 用户角色：user / admin
	Userrole string `json:"userrole"`
	// 创建时间
	Createtime time.Time `json:"createtime"`
	// 更新时间
	Updatetime time.Time `json:"updatetime"`
}

// 注册用户
type CreateUserParamsJSON struct {
	Username     sql.NullString `json:"username"`
	Useraccount  string         `json:"useraccount"`
	Useravatar   sql.NullString `json:"useravatar"`
	Gender       int32          `json:"gender"`
	Userrole     string         `json:"userrole"`
	Userpassword string         `json:"userpassword"`
	Accesskey    string         `json:"accesskey"`
	Secretkey    string         `json:"secretkey"`
}

func ConvertToNormalUser(u *dbsq.User) *ShowUserJSON {
	nu := &ShowUserJSON{
		ID:          u.ID,
		Useraccount: u.Useraccount,
		Userrole:    u.Userrole,
		Createtime:  u.Createtime,
		Updatetime:  u.Updatetime,
	}

	if u.Username.Valid {
		nu.Username = u.Username.String
	}

	if u.Useravatar.Valid {
		nu.Useravatar = u.Useravatar.String
	}

	if u.Gender.Valid {
		nu.Gender = u.Gender.Int32
	}

	return nu
}

func ConvertToCreateUserParamsJSON(jsonParams *CreateUserParamsJSON) *dbsq.CreateUserParams {
	return &dbsq.CreateUserParams{
		Username:    jsonParams.Username,
		Useraccount: jsonParams.Useraccount,
		Useravatar:  jsonParams.Useravatar,
		Gender: sql.NullInt32{
			Int32: jsonParams.Gender,
			Valid: true,
		},
		Userrole:     jsonParams.Userrole,
		Userpassword: jsonParams.Userpassword,
		Accesskey:    jsonParams.Accesskey,
		Secretkey:    jsonParams.Secretkey,
	}
}
