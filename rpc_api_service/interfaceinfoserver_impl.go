package rpcapiservice

import (
	"context"
	glog "xj/xapi-backend/g_log"
	"xj/xapi-backend/models"
	"xj/xapi-backend/rpc_api"
	"xj/xapi-backend/service"
)

type IntefaceInfoServerImpl struct {
	rpc_api.UnimplementedIntefaceInfoServer
}

func (s *IntefaceInfoServerImpl) GetInterfaceInfoById(ctx context.Context, in *rpc_api.GetInterfaceInfoByIdReq) (*rpc_api.GetInterfaceInfoByIdResp, error) {
	glog.Log.Infof("Dubbo-go GetInterfaceInfoById: InterfaceId = %d", in.InterfaceId)
	data, err := service.GetInterfaceInfoById(in.InterfaceId)
	if err != nil {
		return nil, err
	}
	return Convert2GetInterfaceInfoByIdResp(data), nil
}

// 定义一个函数将 ValidInterfaceInfo 结构体转换为 GetInterfaceInfoByIdResp 消息
func Convert2GetInterfaceInfoByIdResp(info *models.ValidInterfaceInfo) *rpc_api.GetInterfaceInfoByIdResp {
	createTime := ConvertTimeToTimestamp(info.Createtime)
	updateTime := ConvertTimeToTimestamp(info.Updatetime)

	// 创建 GetInterfaceInfoByIdResp 消息并赋值字段
	return &rpc_api.GetInterfaceInfoByIdResp{
		Id:             info.ID,
		Name:           info.Name,
		Description:    info.Description,
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
