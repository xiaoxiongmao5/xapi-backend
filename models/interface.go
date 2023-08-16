package models

import (
	"time"
	"xj/xapi-backend/dbsq"
)

// type CreateInterfaceParams struct {
// 	Name           string `json:"name"`
// 	Description    string `json:"description"`
// 	Url            string `json:"url"`
// 	Requestparams  string `json:"requestparams"`
// 	Requestheader  string `json:"requestheader"`
// 	Responseheader string `json:"responseheader"`
// 	Method         string `json:"method"`
// 	Userid         int64  `json:"userid"`
// }

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
	// 请求参数
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
		Requestparams:  i.Requestparams,
		Requestheader:  i.Requestheader.String,  // Get the string value from sql.NullString
		Responseheader: i.Responseheader.String, // Get the string value from sql.NullString
		Status:         i.Status,
		Method:         i.Method,
		Userid:         i.Userid,
		Createtime:     i.Createtime,
		Updatetime:     i.Updatetime,
	}
}
