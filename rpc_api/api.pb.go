// protoc --go_out=. --go-triple_out=. ./api.proto
// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: api.proto

package rpc_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInvokeUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey string `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
}

func (x *GetInvokeUserReq) Reset() {
	*x = GetInvokeUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvokeUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvokeUserReq) ProtoMessage() {}

func (x *GetInvokeUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvokeUserReq.ProtoReflect.Descriptor instead.
func (*GetInvokeUserReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvokeUserReq) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

type GetInvokeUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Useraccount  string                 `protobuf:"bytes,2,opt,name=useraccount,proto3" json:"useraccount,omitempty"`
	Userpassword string                 `protobuf:"bytes,3,opt,name=userpassword,proto3" json:"userpassword,omitempty"`
	Userrole     string                 `protobuf:"bytes,4,opt,name=userrole,proto3" json:"userrole,omitempty"`
	Username     string                 `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Useravatar   string                 `protobuf:"bytes,6,opt,name=useravatar,proto3" json:"useravatar,omitempty"`
	Gender       int32                  `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Accesskey    string                 `protobuf:"bytes,8,opt,name=accesskey,proto3" json:"accesskey,omitempty"`
	Secretkey    string                 `protobuf:"bytes,9,opt,name=secretkey,proto3" json:"secretkey,omitempty"`
	Createtime   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=createtime,proto3" json:"createtime,omitempty"`
	Updatetime   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
	Isdelete     int32                  `protobuf:"varint,12,opt,name=isdelete,proto3" json:"isdelete,omitempty"`
}

func (x *GetInvokeUserResp) Reset() {
	*x = GetInvokeUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvokeUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvokeUserResp) ProtoMessage() {}

func (x *GetInvokeUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvokeUserResp.ProtoReflect.Descriptor instead.
func (*GetInvokeUserResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvokeUserResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetInvokeUserResp) GetUseraccount() string {
	if x != nil {
		return x.Useraccount
	}
	return ""
}

func (x *GetInvokeUserResp) GetUserpassword() string {
	if x != nil {
		return x.Userpassword
	}
	return ""
}

func (x *GetInvokeUserResp) GetUserrole() string {
	if x != nil {
		return x.Userrole
	}
	return ""
}

func (x *GetInvokeUserResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetInvokeUserResp) GetUseravatar() string {
	if x != nil {
		return x.Useravatar
	}
	return ""
}

func (x *GetInvokeUserResp) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *GetInvokeUserResp) GetAccesskey() string {
	if x != nil {
		return x.Accesskey
	}
	return ""
}

func (x *GetInvokeUserResp) GetSecretkey() string {
	if x != nil {
		return x.Secretkey
	}
	return ""
}

func (x *GetInvokeUserResp) GetCreatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Createtime
	}
	return nil
}

func (x *GetInvokeUserResp) GetUpdatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Updatetime
	}
	return nil
}

func (x *GetInvokeUserResp) GetIsdelete() int32 {
	if x != nil {
		return x.Isdelete
	}
	return 0
}

type InvokeCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceId int64 `protobuf:"varint,1,opt,name=interfaceId,proto3" json:"interfaceId,omitempty"`
	UserId      int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *InvokeCountReq) Reset() {
	*x = InvokeCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeCountReq) ProtoMessage() {}

func (x *InvokeCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeCountReq.ProtoReflect.Descriptor instead.
func (*InvokeCountReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *InvokeCountReq) GetInterfaceId() int64 {
	if x != nil {
		return x.InterfaceId
	}
	return 0
}

func (x *InvokeCountReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type InvokeCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InvokeCountResp) Reset() {
	*x = InvokeCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeCountResp) ProtoMessage() {}

func (x *InvokeCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeCountResp.ProtoReflect.Descriptor instead.
func (*InvokeCountResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *InvokeCountResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type GetFullUserInterfaceInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceId int64 `protobuf:"varint,1,opt,name=interfaceId,proto3" json:"interfaceId,omitempty"`
	UserId      int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetFullUserInterfaceInfoReq) Reset() {
	*x = GetFullUserInterfaceInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullUserInterfaceInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullUserInterfaceInfoReq) ProtoMessage() {}

func (x *GetFullUserInterfaceInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullUserInterfaceInfoReq.ProtoReflect.Descriptor instead.
func (*GetFullUserInterfaceInfoReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetFullUserInterfaceInfoReq) GetInterfaceId() int64 {
	if x != nil {
		return x.InterfaceId
	}
	return 0
}

func (x *GetFullUserInterfaceInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFullUserInterfaceInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Totalnum       int32                  `protobuf:"varint,4,opt,name=totalnum,proto3" json:"totalnum,omitempty"`
	Leftnum        int32                  `protobuf:"varint,5,opt,name=leftnum,proto3" json:"leftnum,omitempty"`
	Banstatus      int32                  `protobuf:"varint,6,opt,name=banstatus,proto3" json:"banstatus,omitempty"`
	Host           string                 `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Url            string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Requestparams  string                 `protobuf:"bytes,9,opt,name=requestparams,proto3" json:"requestparams,omitempty"`
	Requestheader  string                 `protobuf:"bytes,10,opt,name=requestheader,proto3" json:"requestheader,omitempty"`
	Responseheader string                 `protobuf:"bytes,11,opt,name=responseheader,proto3" json:"responseheader,omitempty"`
	Status         int32                  `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	Method         string                 `protobuf:"bytes,13,opt,name=method,proto3" json:"method,omitempty"`
	Userid         int64                  `protobuf:"varint,14,opt,name=userid,proto3" json:"userid,omitempty"`
	Createtime     *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=createtime,proto3" json:"createtime,omitempty"`
	Updatetime     *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
}

func (x *GetFullUserInterfaceInfoResp) Reset() {
	*x = GetFullUserInterfaceInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullUserInterfaceInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullUserInterfaceInfoResp) ProtoMessage() {}

func (x *GetFullUserInterfaceInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullUserInterfaceInfoResp.ProtoReflect.Descriptor instead.
func (*GetFullUserInterfaceInfoResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetFullUserInterfaceInfoResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetTotalnum() int32 {
	if x != nil {
		return x.Totalnum
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetLeftnum() int32 {
	if x != nil {
		return x.Leftnum
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetBanstatus() int32 {
	if x != nil {
		return x.Banstatus
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetRequestparams() string {
	if x != nil {
		return x.Requestparams
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetRequestheader() string {
	if x != nil {
		return x.Requestheader
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetResponseheader() string {
	if x != nil {
		return x.Responseheader
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GetFullUserInterfaceInfoResp) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *GetFullUserInterfaceInfoResp) GetCreatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Createtime
	}
	return nil
}

func (x *GetFullUserInterfaceInfoResp) GetUpdatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Updatetime
	}
	return nil
}

type GetInterfaceInfoByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceId int64 `protobuf:"varint,1,opt,name=interfaceId,proto3" json:"interfaceId,omitempty"`
}

func (x *GetInterfaceInfoByIdReq) Reset() {
	*x = GetInterfaceInfoByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterfaceInfoByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterfaceInfoByIdReq) ProtoMessage() {}

func (x *GetInterfaceInfoByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterfaceInfoByIdReq.ProtoReflect.Descriptor instead.
func (*GetInterfaceInfoByIdReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetInterfaceInfoByIdReq) GetInterfaceId() int64 {
	if x != nil {
		return x.InterfaceId
	}
	return 0
}

type GetInterfaceInfoByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Host           string                 `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Url            string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Requestparams  string                 `protobuf:"bytes,6,opt,name=requestparams,proto3" json:"requestparams,omitempty"`
	Requestheader  string                 `protobuf:"bytes,7,opt,name=requestheader,proto3" json:"requestheader,omitempty"`
	Responseheader string                 `protobuf:"bytes,8,opt,name=responseheader,proto3" json:"responseheader,omitempty"`
	Status         int32                  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Method         string                 `protobuf:"bytes,10,opt,name=method,proto3" json:"method,omitempty"`
	Userid         int64                  `protobuf:"varint,11,opt,name=userid,proto3" json:"userid,omitempty"`
	Createtime     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=createtime,proto3" json:"createtime,omitempty"`
	Updatetime     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updatetime,proto3" json:"updatetime,omitempty"`
}

func (x *GetInterfaceInfoByIdResp) Reset() {
	*x = GetInterfaceInfoByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterfaceInfoByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterfaceInfoByIdResp) ProtoMessage() {}

func (x *GetInterfaceInfoByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterfaceInfoByIdResp.ProtoReflect.Descriptor instead.
func (*GetInterfaceInfoByIdResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetInterfaceInfoByIdResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetInterfaceInfoByIdResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetRequestparams() string {
	if x != nil {
		return x.Requestparams
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetRequestheader() string {
	if x != nil {
		return x.Requestheader
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetResponseheader() string {
	if x != nil {
		return x.Responseheader
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetInterfaceInfoByIdResp) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GetInterfaceInfoByIdResp) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *GetInterfaceInfoByIdResp) GetCreatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Createtime
	}
	return nil
}

func (x *GetInterfaceInfoByIdResp) GetUpdatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Updatetime
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x22, 0xa9, 0x03, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x57, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x92, 0x04, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x6e, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x6e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xba, 0x03, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x24,
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x32, 0x54, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x32, 0xc1, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a,
	0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x72,
	0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x69, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x32, 0x6d, 0x0a, 0x0c,
	0x49, 0x6e, 0x74, 0x65, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x3b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(*GetInvokeUserReq)(nil),             // 0: rpc_api.GetInvokeUserReq
	(*GetInvokeUserResp)(nil),            // 1: rpc_api.GetInvokeUserResp
	(*InvokeCountReq)(nil),               // 2: rpc_api.InvokeCountReq
	(*InvokeCountResp)(nil),              // 3: rpc_api.InvokeCountResp
	(*GetFullUserInterfaceInfoReq)(nil),  // 4: rpc_api.GetFullUserInterfaceInfoReq
	(*GetFullUserInterfaceInfoResp)(nil), // 5: rpc_api.GetFullUserInterfaceInfoResp
	(*GetInterfaceInfoByIdReq)(nil),      // 6: rpc_api.GetInterfaceInfoByIdReq
	(*GetInterfaceInfoByIdResp)(nil),     // 7: rpc_api.GetInterfaceInfoByIdResp
	(*timestamppb.Timestamp)(nil),        // 8: google.protobuf.Timestamp
}
var file_api_proto_depIdxs = []int32{
	8,  // 0: rpc_api.GetInvokeUserResp.createtime:type_name -> google.protobuf.Timestamp
	8,  // 1: rpc_api.GetInvokeUserResp.updatetime:type_name -> google.protobuf.Timestamp
	8,  // 2: rpc_api.GetFullUserInterfaceInfoResp.createtime:type_name -> google.protobuf.Timestamp
	8,  // 3: rpc_api.GetFullUserInterfaceInfoResp.updatetime:type_name -> google.protobuf.Timestamp
	8,  // 4: rpc_api.GetInterfaceInfoByIdResp.createtime:type_name -> google.protobuf.Timestamp
	8,  // 5: rpc_api.GetInterfaceInfoByIdResp.updatetime:type_name -> google.protobuf.Timestamp
	0,  // 6: rpc_api.UserInfo.GetInvokeUser:input_type -> rpc_api.GetInvokeUserReq
	2,  // 7: rpc_api.UserIntefaceInfo.InvokeCount:input_type -> rpc_api.InvokeCountReq
	4,  // 8: rpc_api.UserIntefaceInfo.GetFullUserInterfaceInfo:input_type -> rpc_api.GetFullUserInterfaceInfoReq
	6,  // 9: rpc_api.IntefaceInfo.GetInterfaceInfoById:input_type -> rpc_api.GetInterfaceInfoByIdReq
	1,  // 10: rpc_api.UserInfo.GetInvokeUser:output_type -> rpc_api.GetInvokeUserResp
	3,  // 11: rpc_api.UserIntefaceInfo.InvokeCount:output_type -> rpc_api.InvokeCountResp
	5,  // 12: rpc_api.UserIntefaceInfo.GetFullUserInterfaceInfo:output_type -> rpc_api.GetFullUserInterfaceInfoResp
	7,  // 13: rpc_api.IntefaceInfo.GetInterfaceInfoById:output_type -> rpc_api.GetInterfaceInfoByIdResp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvokeUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvokeUserResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeCountReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeCountResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullUserInterfaceInfoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullUserInterfaceInfoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterfaceInfoByIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterfaceInfoByIdResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
