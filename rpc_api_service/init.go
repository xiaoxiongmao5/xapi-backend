package rpcapiservice

import (
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	config.SetProviderService(&UserInfoServerImpl{})
	config.SetProviderService(&IntefaceInfoServerImpl{})
	config.SetProviderService(&UserIntefaceInfoServerImpl{})
}

func ConvertTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}
func ConvertTimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
