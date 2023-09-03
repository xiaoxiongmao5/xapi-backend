// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v4.24.2
// source: api.proto

package rpc_api

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// UserIntefaceInfoClient is the client API for UserIntefaceInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserIntefaceInfoClient interface {
	// 调用接口统计
	InvokeCount(ctx context.Context, in *InvokeCountReq, opts ...grpc_go.CallOption) (*InvokeCountResp, common.ErrorWithAttachment)
}

type userIntefaceInfoClient struct {
	cc *triple.TripleConn
}

type UserIntefaceInfoClientImpl struct {
	InvokeCount func(ctx context.Context, in *InvokeCountReq) (*InvokeCountResp, error)
}

func (c *UserIntefaceInfoClientImpl) GetDubboStub(cc *triple.TripleConn) UserIntefaceInfoClient {
	return NewUserIntefaceInfoClient(cc)
}

func (c *UserIntefaceInfoClientImpl) XXX_InterfaceName() string {
	return "rpc_api.UserIntefaceInfo"
}

func NewUserIntefaceInfoClient(cc *triple.TripleConn) UserIntefaceInfoClient {
	return &userIntefaceInfoClient{cc}
}

func (c *userIntefaceInfoClient) InvokeCount(ctx context.Context, in *InvokeCountReq, opts ...grpc_go.CallOption) (*InvokeCountResp, common.ErrorWithAttachment) {
	out := new(InvokeCountResp)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/InvokeCount", in, out)
}

// UserIntefaceInfoServer is the server API for UserIntefaceInfo service.
// All implementations must embed UnimplementedUserIntefaceInfoServer
// for forward compatibility
type UserIntefaceInfoServer interface {
	// 调用接口统计
	InvokeCount(context.Context, *InvokeCountReq) (*InvokeCountResp, error)
	mustEmbedUnimplementedUserIntefaceInfoServer()
}

// UnimplementedUserIntefaceInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserIntefaceInfoServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedUserIntefaceInfoServer) InvokeCount(context.Context, *InvokeCountReq) (*InvokeCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeCount not implemented")
}
func (s *UnimplementedUserIntefaceInfoServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedUserIntefaceInfoServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedUserIntefaceInfoServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &UserIntefaceInfo_ServiceDesc
}
func (s *UnimplementedUserIntefaceInfoServer) XXX_InterfaceName() string {
	return "rpc_api.UserIntefaceInfo"
}

func (UnimplementedUserIntefaceInfoServer) mustEmbedUnimplementedUserIntefaceInfoServer() {}

// UnsafeUserIntefaceInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserIntefaceInfoServer will
// result in compilation errors.
type UnsafeUserIntefaceInfoServer interface {
	mustEmbedUnimplementedUserIntefaceInfoServer()
}

func RegisterUserIntefaceInfoServer(s grpc_go.ServiceRegistrar, srv UserIntefaceInfoServer) {
	s.RegisterService(&UserIntefaceInfo_ServiceDesc, srv)
}

func _UserIntefaceInfo_InvokeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("InvokeCount", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// UserIntefaceInfo_ServiceDesc is the grpc_go.ServiceDesc for UserIntefaceInfo service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserIntefaceInfo_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "rpc_api.UserIntefaceInfo",
	HandlerType: (*UserIntefaceInfoServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "InvokeCount",
			Handler:    _UserIntefaceInfo_InvokeCount_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "api.proto",
}

// IntefaceInfoClient is the client API for IntefaceInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntefaceInfoClient interface {
	// 从数据库中查询接口是否存在（请求路径、请求方法、请求参数）
	GetInterfaceInfo(ctx context.Context, in *GetInterfaceInfoReq, opts ...grpc_go.CallOption) (*GetInterfaceInfoResp, common.ErrorWithAttachment)
}

type intefaceInfoClient struct {
	cc *triple.TripleConn
}

type IntefaceInfoClientImpl struct {
	GetInterfaceInfo func(ctx context.Context, in *GetInterfaceInfoReq) (*GetInterfaceInfoResp, error)
}

func (c *IntefaceInfoClientImpl) GetDubboStub(cc *triple.TripleConn) IntefaceInfoClient {
	return NewIntefaceInfoClient(cc)
}

func (c *IntefaceInfoClientImpl) XXX_InterfaceName() string {
	return "rpc_api.IntefaceInfo"
}

func NewIntefaceInfoClient(cc *triple.TripleConn) IntefaceInfoClient {
	return &intefaceInfoClient{cc}
}

func (c *intefaceInfoClient) GetInterfaceInfo(ctx context.Context, in *GetInterfaceInfoReq, opts ...grpc_go.CallOption) (*GetInterfaceInfoResp, common.ErrorWithAttachment) {
	out := new(GetInterfaceInfoResp)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetInterfaceInfo", in, out)
}

// IntefaceInfoServer is the server API for IntefaceInfo service.
// All implementations must embed UnimplementedIntefaceInfoServer
// for forward compatibility
type IntefaceInfoServer interface {
	// 从数据库中查询接口是否存在（请求路径、请求方法、请求参数）
	GetInterfaceInfo(context.Context, *GetInterfaceInfoReq) (*GetInterfaceInfoResp, error)
	mustEmbedUnimplementedIntefaceInfoServer()
}

// UnimplementedIntefaceInfoServer must be embedded to have forward compatible implementations.
type UnimplementedIntefaceInfoServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedIntefaceInfoServer) GetInterfaceInfo(context.Context, *GetInterfaceInfoReq) (*GetInterfaceInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterfaceInfo not implemented")
}
func (s *UnimplementedIntefaceInfoServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedIntefaceInfoServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedIntefaceInfoServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &IntefaceInfo_ServiceDesc
}
func (s *UnimplementedIntefaceInfoServer) XXX_InterfaceName() string {
	return "rpc_api.IntefaceInfo"
}

func (UnimplementedIntefaceInfoServer) mustEmbedUnimplementedIntefaceInfoServer() {}

// UnsafeIntefaceInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntefaceInfoServer will
// result in compilation errors.
type UnsafeIntefaceInfoServer interface {
	mustEmbedUnimplementedIntefaceInfoServer()
}

func RegisterIntefaceInfoServer(s grpc_go.ServiceRegistrar, srv IntefaceInfoServer) {
	s.RegisterService(&IntefaceInfo_ServiceDesc, srv)
}

func _IntefaceInfo_GetInterfaceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInterfaceInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetInterfaceInfo", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// IntefaceInfo_ServiceDesc is the grpc_go.ServiceDesc for IntefaceInfo service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntefaceInfo_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "rpc_api.IntefaceInfo",
	HandlerType: (*IntefaceInfoServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetInterfaceInfo",
			Handler:    _IntefaceInfo_GetInterfaceInfo_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "api.proto",
}

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	// 数据库中查询是否已分配给用户秘钥（accessKey）
	GetInvokeUser(ctx context.Context, in *GetInvokeUserReq, opts ...grpc_go.CallOption) (*GetInvokeUserResp, common.ErrorWithAttachment)
}

type userInfoClient struct {
	cc *triple.TripleConn
}

type UserInfoClientImpl struct {
	GetInvokeUser func(ctx context.Context, in *GetInvokeUserReq) (*GetInvokeUserResp, error)
}

func (c *UserInfoClientImpl) GetDubboStub(cc *triple.TripleConn) UserInfoClient {
	return NewUserInfoClient(cc)
}

func (c *UserInfoClientImpl) XXX_InterfaceName() string {
	return "rpc_api.UserInfo"
}

func NewUserInfoClient(cc *triple.TripleConn) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetInvokeUser(ctx context.Context, in *GetInvokeUserReq, opts ...grpc_go.CallOption) (*GetInvokeUserResp, common.ErrorWithAttachment) {
	out := new(GetInvokeUserResp)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetInvokeUser", in, out)
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	// 数据库中查询是否已分配给用户秘钥（accessKey）
	GetInvokeUser(context.Context, *GetInvokeUserReq) (*GetInvokeUserResp, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedUserInfoServer) GetInvokeUser(context.Context, *GetInvokeUserReq) (*GetInvokeUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvokeUser not implemented")
}
func (s *UnimplementedUserInfoServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedUserInfoServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedUserInfoServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &UserInfo_ServiceDesc
}
func (s *UnimplementedUserInfoServer) XXX_InterfaceName() string {
	return "rpc_api.UserInfo"
}

func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc_go.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_GetInvokeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvokeUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetInvokeUser", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc_go.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "rpc_api.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetInvokeUser",
			Handler:    _UserInfo_GetInvokeUser_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "api.proto",
}
