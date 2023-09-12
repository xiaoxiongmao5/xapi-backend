package rpcapiservice

import (
	"context"

	"xj/xapi-backend/dbsq"
	glog "xj/xapi-backend/g_log"
	"xj/xapi-backend/rpc_api"
	"xj/xapi-backend/service"
)

type UserInfoServerImpl struct {
	rpc_api.UnimplementedUserInfoServer
}

func (s *UserInfoServerImpl) GetInvokeUser(ctx context.Context, in *rpc_api.GetInvokeUserReq) (*rpc_api.GetInvokeUserResp, error) {
	glog.Log.Infof("Dubbo-go GetInvokeUser AccessKey = %s", in.AccessKey)
	data, err := service.GetUserInfoByAccessKey(in.AccessKey)
	if err != nil {
		return nil, err
	}
	return ConvertUserToGetInvokeUserResp(data)
}

// 定义一个函数将 User 结构体转换为 GetInvokeUserResp 消息
func ConvertUserToGetInvokeUserResp(user *dbsq.User) (*rpc_api.GetInvokeUserResp, error) {
	createTime := ConvertTimeToTimestamp(user.Createtime)

	updateTime := ConvertTimeToTimestamp(user.Updatetime)

	// 创建 GetInvokeUserResp 消息并赋值字段
	resp := &rpc_api.GetInvokeUserResp{
		Id:           user.ID,
		Useraccount:  user.Useraccount,
		Userpassword: user.Userpassword,
		Userrole:     user.Userrole,
		Username:     user.Username.String, // 使用 sql.NullString 的 String 方法获取值
		Useravatar:   user.Useravatar.String,
		Gender:       user.Gender.Int32, // 使用 sql.NullInt32 的 Int32 方法获取值
		Accesskey:    user.Accesskey,
		Secretkey:    user.Secretkey,
		Createtime:   createTime,
		Updatetime:   updateTime,
		Isdelete:     user.Isdelete,
	}

	return resp, nil
}
