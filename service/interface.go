package service

import (
	"context"
	"database/sql"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/enums"
	"xj/xapi-backend/models"
)

// 根据接口ID 获得接口信息
func GetInterfaceInfo(id int64) (*models.ValidXapiInterfaceInfo, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	data, err := q.GetInterfaceInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	retdata := models.ConvertToValidXapiInterfaceInfo(data)
	return retdata, nil
}

// 注册一条接口
func CreateInterface(param *models.CreateInterfaceParams) (sql.Result, error) {
	nparam := models.ConvertToCreateInterfaceParams(param)
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.CreateInterface(ctx, nparam)
}

// 更新接口信息
func UpdateInterface(param *models.UpdateInterfaceParams) error {
	nparam := models.ConvertToUpdateInterfaceParams(param)
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.UpdateInterface(ctx, nparam)
}

// 删除一条接口
func DeleteInterface(id int64) error {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.DeleteInterface(ctx, id)
}

func ConvertSliceToValidXapiInterfaceInfo(slice []*dbsq.XapiInterfaceInfo) []*models.ValidXapiInterfaceInfo {
	result := make([]*models.ValidXapiInterfaceInfo, len(slice))
	for i, item := range slice {
		result[i] = models.ConvertToValidXapiInterfaceInfo(item)
	}
	return result
}

// 获得所有接口列表
func ListInterfaces() ([]*models.ValidXapiInterfaceInfo, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	res, error := q.ListInterfaces(ctx)
	if error != nil {
		return nil, error
	}
	return ConvertSliceToValidXapiInterfaceInfo(res), nil
}

// 发布接口
func OnlineInterfaceStatus(id int64) error {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	param := &dbsq.UpdateInterfaceStatusParams{
		Status: int32(enums.InterfaceStatusOnline),
		ID:     id,
	}
	return q.UpdateInterfaceStatus(ctx, param)
}

// 下线接口
func OfflineInterfaceStatus(id int64) error {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	param := &dbsq.UpdateInterfaceStatusParams{
		Status: int32(enums.InterfaceStatusOffline),
		ID:     id,
	}
	return q.UpdateInterfaceStatus(ctx, param)
}
