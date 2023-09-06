package rpcapiservice

import (
	"context"
	"xj/xapi-backend/models"
	"xj/xapi-backend/rpc_api"
	"xj/xapi-backend/service"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type IntefaceInfoServerImpl struct {
	rpc_api.UnimplementedIntefaceInfoServer
}

func (s *IntefaceInfoServerImpl) GetInterfaceInfoById(ctx context.Context, in *rpc_api.GetInterfaceInfoByIdReq) (*rpc_api.GetInterfaceInfoByIdResp, error) {
	logger.Infof("Dubbo-go GetInterfaceInfoById: InterfaceId = %d\n", in.InterfaceId)
	data, err := service.GetInterfaceInfoById(in.InterfaceId)
	if err != nil {
		return nil, err
	}
	return ConvertValidXapiInterfaceInfoToGetInterfaceInfoByIdResp(data)
}

// 定义一个函数将 ValidXapiInterfaceInfo 结构体转换为 GetInterfaceInfoByIdResp 消息
func ConvertValidXapiInterfaceInfoToGetInterfaceInfoByIdResp(info *models.ValidXapiInterfaceInfo) (*rpc_api.GetInterfaceInfoByIdResp, error) {
	createTime := ConvertTimeToTimestamp(info.Createtime)
	updateTime := ConvertTimeToTimestamp(info.Updatetime)

	// 创建 GetInterfaceInfoByIdResp 消息并赋值字段
	resp := &rpc_api.GetInterfaceInfoByIdResp{
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

	return resp, nil
}
