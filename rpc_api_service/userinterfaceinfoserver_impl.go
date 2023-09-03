package rpcapiservice

import (
	"context"
	"xj/xapi-backend/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
)

type UserIntefaceInfoServerImpl struct {
	rpc_api.UnimplementedUserIntefaceInfoServer
}

func (s *UserIntefaceInfoServerImpl) InvokeCount(ctx context.Context, in *rpc_api.InvokeCountReq) (*rpc_api.InvokeCountResp, error) {
	logger.Infof("Dubbo-go InvokeCount InterfaceId = %s UserId = %s\n", in.InterfaceId, in.UserId)
	return &rpc_api.InvokeCountResp{
		Result: true,
	}, nil
}
