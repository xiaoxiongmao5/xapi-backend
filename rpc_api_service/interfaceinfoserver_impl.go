package rpcapiservice

import (
	"context"
	"xj/xapi-backend/models"
	"xj/xapi-backend/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type IntefaceInfoServerImpl struct {
	rpc_api.UnimplementedIntefaceInfoServer
}

func (s *IntefaceInfoServerImpl) GetInterfaceInfo(ctx context.Context, in *rpc_api.GetInterfaceInfoReq) (*rpc_api.GetInterfaceInfoResp, error) {
	logger.Infof("Dubbo-go GetInterfaceInfo Path = %s Method = %s\n", in.Path, in.Method)
	return &rpc_api.GetInterfaceInfoResp{}, nil
}

// 定义一个函数将 ValidXapiInterfaceInfo 结构体转换为 GetInterfaceInfoResp 消息
func ConvertValidXapiInterfaceInfoToGetInterfaceInfoResp(info *models.ValidXapiInterfaceInfo) (*rpc_api.GetInterfaceInfoResp, error) {
	createTime := ConvertTimeToTimestamp(info.Createtime)

	updateTime := ConvertTimeToTimestamp(info.Updatetime)

	// 创建 GetInterfaceInfoResp 消息并赋值字段
	resp := &rpc_api.GetInterfaceInfoResp{
		Id:             info.ID,
		Name:           info.Name,
		Description:    info.Description,
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

	return resp, nil
}
