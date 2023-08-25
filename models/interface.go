package models

import (
	"database/sql"
	"time"
	"xj/xapi-backend/dbsq"
)

type CreateInterfaceParams struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Requestparams  string `json:"requestparams"`
	Requestheader  string `json:"requestheader"`
	Responseheader string `json:"responseheader"`
	Method         string `json:"method"`
	Userid         int64  `json:"userid"`
}

type UpdateInterfaceParams struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Requestparams  string `json:"requestparams"`
	Requestheader  string `json:"requestheader"`
	Responseheader string `json:"responseheader"`
	Method         string `json:"method"`
	Userid         int64  `json:"userid"`
	ID             int64  `json:"id"`
}

// 接口信息
type ValidXapiInterfaceInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 描述
	Description string `json:"description"`
	// 接口地址
	Url string `json:"url"`
	/** 请求参数
	[
		{"name":"xxx", "type":"string"}
	]
	*/
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

func ConvertToValidXapiInterfaceInfo(i *dbsq.XapiInterfaceInfo) *ValidXapiInterfaceInfo {
	return &ValidXapiInterfaceInfo{
		ID:             i.ID,
		Name:           i.Name,
		Description:    i.Description.String, // Get the string value from sql.NullString
		Url:            i.Url,
		Requestparams:  i.Requestparams.String,
		Requestheader:  i.Requestheader.String,  // Get the string value from sql.NullString
		Responseheader: i.Responseheader.String, // Get the string value from sql.NullString
		Status:         i.Status,
		Method:         i.Method,
		Userid:         i.Userid,
		Createtime:     i.Createtime,
		Updatetime:     i.Updatetime,
	}
}

func ConvertToCreateInterfaceParams(jsonParams *CreateInterfaceParams) *dbsq.CreateInterfaceParams {
	return &dbsq.CreateInterfaceParams{
		Name: jsonParams.Name,
		Description: sql.NullString{
			String: jsonParams.Description,
			Valid:  true,
		},
		Url: jsonParams.Url,
		Requestparams: sql.NullString{
			String: jsonParams.Requestparams,
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

func ConvertToUpdateInterfaceParams(jsonParams *UpdateInterfaceParams) *dbsq.UpdateInterfaceParams {
	return &dbsq.UpdateInterfaceParams{
		Name: jsonParams.Name,
		Description: sql.NullString{
			String: jsonParams.Description,
			Valid:  true,
		},
		Url: jsonParams.Url,
		Requestparams: sql.NullString{
			String: jsonParams.Requestparams,
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
