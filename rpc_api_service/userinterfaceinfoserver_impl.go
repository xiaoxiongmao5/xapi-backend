package rpcapiservice

import (
	"context"
	"xj/xapi-backend/rpc_api"
	"xj/xapi-backend/service"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type UserIntefaceInfoServerImpl struct {
	rpc_api.UnimplementedUserIntefaceInfoServer
}

func (s *UserIntefaceInfoServerImpl) InvokeCount(ctx context.Context, in *rpc_api.InvokeCountReq) (*rpc_api.InvokeCountResp, error) {
	logger.Infof("Dubbo-go InvokeCount InterfaceId = %s UserId = %s\n", in.InterfaceId, in.UserId)
	data, err := service.InvokeCount(in.InterfaceId, in.UserId)
	if err != nil {
		return &rpc_api.InvokeCountResp{
			Result: false,
		}, err
	}

	if num, err := data.RowsAffected(); err != nil || num == 0 {
		return &rpc_api.InvokeCountResp{
			Result: false,
		}, err
	}

	return &rpc_api.InvokeCountResp{
		Result: true,
	}, nil
}
