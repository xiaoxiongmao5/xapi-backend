package service

import (
	"context"
	"database/sql"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/enums"
	"xj/xapi-backend/models"
	"xj/xapi-backend/utils"
)

// 根据接口ID 获得接口信息
func GetInterfaceInfoById(id int64) (*models.ValidInterfaceInfo, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	data, err := q.GetInterfaceInfoById(ctx, id)
	if err != nil {
		return nil, err
	}
	retdata := models.Convert2ValidXapiInterfaceInfo(data)
	return retdata, nil
}

// 根据 host+url+method 获得接口信息
func GetInterfaceInfoByUniFullApi(host, url, method string) (*models.ValidInterfaceInfo, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	param := &dbsq.GetInterfaceInfoByUniFullApiParams{
		Host:   host,
		Url:    url,
		Method: method,
	}
	data, err := q.GetInterfaceInfoByUniFullApi(ctx, param)
	if err != nil {
		return nil, err
	}
	retdata := models.Convert2ValidXapiInterfaceInfo(data)
	return retdata, nil
}

// 注册一条接口
func CreateInterface(param *models.CreateInterfaceParams) (sql.Result, error) {
	nparam := models.Convert2CreateInterfaceParams(param)
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.CreateInterface(ctx, nparam)
}

// 更新接口信息
func UpdateInterface(param *models.UpdateInterfaceParams) error {
	nparam := models.Convert2UpdateInterfaceParams(param)
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.UpdateInterface(ctx, nparam)
}

// 删除一条接口
func DeleteInterface(id int64) error {
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.DeleteInterface(ctx, id)
}

func ConvertSliceToValidXapiInterfaceInfo(slice []*dbsq.XapiInterfaceInfo) []*models.ValidInterfaceInfo {
	result := make([]*models.ValidInterfaceInfo, len(slice))
	for i, item := range slice {
		result[i] = models.Convert2ValidXapiInterfaceInfo(item)
	}
	return result
}

// 获得所有接口列表
func AllListInterfaces() ([]*models.ValidInterfaceInfo, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	res, error := q.ListInterfaces(ctx)
	if error != nil {
		return nil, error
	}
	return ConvertSliceToValidXapiInterfaceInfo(res), nil
}

// 分页获得接口列表
func PageListInterfaces(current, pageSize int) ([]*models.ValidInterfaceInfo, error) {
	limit, offset := utils.CalculateLimitOffset(current, pageSize)
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	res, error := q.ListPageInterfaces(ctx, &dbsq.ListPageInterfacesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if error != nil {
		return nil, error
	}
	return ConvertSliceToValidXapiInterfaceInfo(res), nil
}

// 获得接口列表总条数
func GetInterfaceListCount() (int64, error) {
	conn, err := db.GetConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetInterfaceListCount(ctx)
}

// 分页获得已发布的接口列表（status=1）
func PageListOnlineInterfaces(current, pageSize int) ([]*models.ValidInterfaceInfo, error) {
	limit, offset := utils.CalculateLimitOffset(current, pageSize)
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	res, error := q.ListPageOnlineInterfaces(ctx, &dbsq.ListPageOnlineInterfacesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if error != nil {
		return nil, error
	}
	return ConvertSliceToValidXapiInterfaceInfo(res), nil
}

// 获得已发布的接口列表总数（status=1）
func GetOnlineInterfaceListCount() (int64, error) {
	conn, err := db.GetConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetOnlineInterfaceListCount(ctx)
}

// 发布接口
func OnlineInterfaceStatus(id int64) error {
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.UpdateInterfaceStatus(ctx, &dbsq.UpdateInterfaceStatusParams{
		Status: int32(enums.InterfaceStatusOnline),
		ID:     id,
	})
}

// 下线接口
func OfflineInterfaceStatus(id int64) error {
	conn, err := db.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	param := &dbsq.UpdateInterfaceStatusParams{
		Status: int32(enums.InterfaceStatusOffline),
		ID:     id,
	}
	return q.UpdateInterfaceStatus(ctx, param)
}
