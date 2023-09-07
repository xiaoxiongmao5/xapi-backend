package models

import (
	"time"
	"xj/xapi-backend/dbsq"
)

// 更新用户接口剩余调用次数
type UpdateInvokeLeftCountParams struct {
	Interfaceinfoid int64 `json:"interfaceId"`
	Leftnum         int32 `json:"leftNum"`
}

// 根据用户信息、接口id获取用户接口信息 响应值
type ValidUserInterfaceInfo struct {
	// 主键(接口ID)
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 描述
	Description string `json:"description"`
	// 总调用次数
	Totalnum int32 `json:"totalnum"`
	// 剩余调用次数
	Leftnum   int32 `json:"leftnum"`
	BanStatus int32 `json:"banstatus"`
	// 接口地址
	Host string `json:"host"`
	// 接口地址
	Url string `json:"url"`
	// 请求参数 [{"name":"xxx", "type":"string"}]
	Requestparams string `json:"requestparams"`
	// 请求头
	Requestheader string `json:"requestheader"`
	// 响应头
	Responseheader string `json:"responseheader"`
	// 接口状态（0-关闭，1-开启）
	Status int32 `json:"status"`
	// 请求类型
	Method string `json:"method"`
	// 创建人
	Userid int64 `json:"userid"`
	// 创建时间
	Createtime time.Time `json:"createtime"`
	// 更新时间
	Updatetime time.Time `json:"updatetime"`
}

func Convert2ValidUserInterfaceInfo(i *ValidInterfaceInfo) *ValidUserInterfaceInfo {
	return &ValidUserInterfaceInfo{
		ID:             i.ID,
		Name:           i.Name,
		Description:    i.Description,
		Totalnum:       0,
		Leftnum:        0,
		Host:           i.Host,
		Url:            i.Url,
		Requestparams:  i.Requestparams,
		Requestheader:  i.Requestheader,
		Responseheader: i.Responseheader,
		Status:         i.Status,
		Method:         i.Method,
		Userid:         i.Userid,
		Createtime:     i.Createtime,
		Updatetime:     i.Updatetime,
	}
}

func Convert2ValidUserInterfaceInfoQueryOfByLeftjoin(i *dbsq.GetFullUserInterfaceInfoByUserIdAndInterfaceIdRow) *ValidUserInterfaceInfo {
	return &ValidUserInterfaceInfo{
		ID:             i.ID.Int64,
		Name:           i.Name.String,
		Description:    i.Description.String,
		Totalnum:       i.Totalnum,
		Leftnum:        i.Leftnum,
		BanStatus:      i.BanStatus,
		Host:           i.Host.String,
		Url:            i.Url.String,
		Requestparams:  i.Requestparams.String,
		Requestheader:  i.Requestheader.String,
		Responseheader: i.Responseheader.String,
		Status:         i.Status.Int32,
		Method:         i.Method.String,
		Userid:         i.Userid.Int64,
		Createtime:     i.Createtime.Time,
		Updatetime:     i.Updatetime.Time,
	}
}
