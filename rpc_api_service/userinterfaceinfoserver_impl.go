package rpcapiservice

import (
	"context"
	glog "xj/xapi-backend/g_log"
	"xj/xapi-backend/models"
	"xj/xapi-backend/rpc_api"
	"xj/xapi-backend/service"
)

type UserIntefaceInfoServerImpl struct {
	rpc_api.UnimplementedUserIntefaceInfoServer
}

func (s *UserIntefaceInfoServerImpl) InvokeCount(ctx context.Context, in *rpc_api.InvokeCountReq) (*rpc_api.InvokeCountResp, error) {
	glog.Log.Infof("Dubbo-go InvokeCount InterfaceId = %d UserId = %d", in.InterfaceId, in.UserId)
	data, err := service.InvokeCount(in.InterfaceId, in.UserId)
	if err != nil {
		return &rpc_api.InvokeCountResp{Result: false}, err
	}

	if num, err := data.RowsAffected(); err != nil || num == 0 {
		return &rpc_api.InvokeCountResp{Result: false}, err
	}

	return &rpc_api.InvokeCountResp{Result: true}, nil
}

func (s *UserIntefaceInfoServerImpl) GetFullUserInterfaceInfo(ctx context.Context, in *rpc_api.GetFullUserInterfaceInfoReq) (*rpc_api.GetFullUserInterfaceInfoResp, error) {
	glog.Log.Infof("Dubbo-go GetFullUserInterfaceInfoByUserIdAndInterfaceId InterfaceId = %d UserId = %d", in.InterfaceId, in.UserId)
	data, err := service.GetFullUserInterfaceInfoByUserIdAndInterfaceId(in.InterfaceId, in.UserId)
	if err != nil {
		return nil, err
	}
	return Convert2GetFullUserInterfaceInfoReq(data), nil
}

// 定义一个函数将 ValidUserInterfaceInfo 结构体转换为 GetFullUserInterfaceInfoResp 消息
func Convert2GetFullUserInterfaceInfoReq(info *models.ValidUserInterfaceInfo) *rpc_api.GetFullUserInterfaceInfoResp {
	createTime := ConvertTimeToTimestamp(info.Createtime)
	updateTime := ConvertTimeToTimestamp(info.Updatetime)

	// 创建 GetInterfaceInfoByIdResp 消息并赋值字段
	return &rpc_api.GetFullUserInterfaceInfoResp{
		Id:             info.ID,
		Name:           info.Name,
		Description:    info.Description,
		Totalnum:       info.Totalnum,
		Leftnum:        info.Leftnum,
		Banstatus:      info.BanStatus,
		Host:           info.Host,
		Url:            info.Url,
		Requestparams:  info.Requestparams,
		Requestheader:  info.Requestheader,
		Responseheader: info.Responseheader,
		Status:         info.Status,
		Method:         info.Method,
		Userid:         info.Userid,
		Createtime:     createTime,
		Updatetime:     updateTime,
	}
}
