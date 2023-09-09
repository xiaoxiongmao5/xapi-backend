package models

import (
	"database/sql"
	"time"
	"xj/xapi-backend/dbsq"
)

// 注册接口参数
type CreateInterfaceParams struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Host                 string `json:"host"`
	Url                  string `json:"url"`
	Requestparams        string `json:"requestparams"`
	Requestparamsexample string `json:"requestparamsexample"`
	Requestheader        string `json:"requestheader"`
	Responseheader       string `json:"responseheader"`
	Method               string `json:"method"`
	Userid               int64  `json:"userid"`
}

// 更新接口信息参数
type UpdateInterfaceParams struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Host                 string `json:"host"`
	Url                  string `json:"url"`
	Requestparams        string `json:"requestparams"`
	Requestparamsexample string `json:"requestparamsexample"`
	Requestheader        string `json:"requestheader"`
	Responseheader       string `json:"responseheader"`
	Method               string `json:"method"`
	Userid               int64  `json:"userid"`
	ID                   int64  `json:"id"`
}

// 调用接口参数
type InvokeInterfaceParams struct {
	ID            int64  `json:"id"`
	Requestparams string `json:"requestparams"`
}

// 根据接口id获取接口信息 响应值
type ValidInterfaceInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 描述
	Description string `json:"description"`
	// 接口地址
	Host string `json:"host"`
	// 接口地址
	Url string `json:"url"`
	// 请求参数
	Requestparams string `json:"requestparams"`
	// 请求参数示例	[{"name":"xxx", "type":"string"}]
	Requestparamsexample string `json:"requestparamsexample"`
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

func Convert2ValidXapiInterfaceInfo(i *dbsq.XapiInterfaceInfo) *ValidInterfaceInfo {
	return &ValidInterfaceInfo{
		ID:                   i.ID,
		Name:                 i.Name,
		Description:          i.Description.String, // Get the string value from sql.NullString
		Host:                 i.Host,
		Url:                  i.Url,
		Requestparams:        i.Requestparams.String,
		Requestparamsexample: i.Requestparamsexample.String,
		Requestheader:        i.Requestheader.String,  // Get the string value from sql.NullString
		Responseheader:       i.Responseheader.String, // Get the string value from sql.NullString
		Status:               i.Status,
		Method:               i.Method,
		Userid:               i.Userid,
		Createtime:           i.Createtime,
		Updatetime:           i.Updatetime,
	}
}

func Convert2CreateInterfaceParams(jsonParams *CreateInterfaceParams) *dbsq.CreateInterfaceParams {
	return &dbsq.CreateInterfaceParams{
		Name: jsonParams.Name,
		Description: sql.NullString{
			String: jsonParams.Description,
			Valid:  true,
		},
		Host: jsonParams.Host,
		Url:  jsonParams.Url,
		Requestparams: sql.NullString{
			String: jsonParams.Requestparams,
			Valid:  true,
		},
		Requestparamsexample: sql.NullString{
			String: jsonParams.Requestparamsexample,
			Valid:  true,
		},
		Requestheader: sql.NullString{
			String: jsonParams.Requestheader,
			Valid:  true,
		},
		Responseheader: sql.NullString{
			String: jsonParams.Responseheader,
			Valid:  true,
		},
		Method: jsonParams.Method,
		Userid: jsonParams.Userid,
	}
}

func Convert2UpdateInterfaceParams(jsonParams *UpdateInterfaceParams) *dbsq.UpdateInterfaceParams {
	return &dbsq.UpdateInterfaceParams{
		Name: jsonParams.Name,
		Description: sql.NullString{
			String: jsonParams.Description,
			Valid:  true,
		},
		Host: jsonParams.Host,
		Url:  jsonParams.Url,
		Requestparams: sql.NullString{
			String: jsonParams.Requestparams,
			Valid:  true,
		},
		Requestparamsexample: sql.NullString{
			String: jsonParams.Requestparamsexample,
			Valid:  true,
		},
		Requestheader: sql.NullString{
			String: jsonParams.Requestheader,
			Valid:  true,
		},
		Responseheader: sql.NullString{
			String: jsonParams.Responseheader,
			Valid:  true,
		},
		Method: jsonParams.Method,
		Userid: jsonParams.Userid,
		ID:     jsonParams.ID,
	}
}
