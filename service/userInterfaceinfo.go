package service

import (
	"context"
	"database/sql"
	"errors"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/models"
)

func GetUserInterfaceInfoByUserIdAndInterfaceId(interfaceId, userId int64) (*dbsq.XapiUserInterfaceInfo, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetUserInterfaceInfoByUserIdAndInterfaceId(ctx, &dbsq.GetUserInterfaceInfoByUserIdAndInterfaceIdParams{
		Interfaceinfoid: interfaceId,
		Userid:          userId,
	})
}

func GetFullUserInterfaceInfoByUserIdAndInterfaceId(interfaceId, userId int64) (*models.ValidUserInterfaceInfo, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	data, err := q.GetFullUserInterfaceInfoByUserIdAndInterfaceId(ctx, &dbsq.GetFullUserInterfaceInfoByUserIdAndInterfaceIdParams{
		Interfaceinfoid: interfaceId,
		Userid:          userId,
	})
	if err != nil {
		return nil, err
	}
	return models.Convert2ValidUserInterfaceInfoQueryOfByLeftjoin(data), nil
}

/** 用户调用接口关系-计数变更
 */
func InvokeCount(interfaceId, userId int64) (sql.Result, error) {
	// 判断接口和用户是否存在
	if interfaceId <= 0 || userId <= 0 {
		return nil, errors.New("用户或接口不存在")
	}
	// 计数变更
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	param := &dbsq.InvokeUserInterfaceInfoParams{
		Interfaceinfoid: interfaceId,
		Userid:          userId,
	}
	return q.InvokeUserInterfaceInfo(ctx, param)
}

/** 更新用户接口剩余调用次数
 */
func UpdateInvokeLeftCount(interfaceId, userId int64, leftNum int32) (sql.Result, error) {
	// 查看该用户是否已经开通
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	data, err := q.GetUserInterfaceInfoByUserIdAndInterfaceId(ctx, &dbsq.GetUserInterfaceInfoByUserIdAndInterfaceIdParams{
		Userid:          userId,
		Interfaceinfoid: interfaceId,
	})
	if err != nil || data == nil {
		// 未开通过
		return q.CreateUserInterfaceInfoWithLeftNum(ctx, &dbsq.CreateUserInterfaceInfoWithLeftNumParams{
			Userid:          userId,
			Interfaceinfoid: interfaceId,
			Leftnum:         leftNum,
		})
	}
	// 开通过，增加剩余次数
	return nil, q.UpdateUserInterfaceInfoLeftNum(ctx, &dbsq.UpdateUserInterfaceInfoLeftNumParams{
		Userid:          userId,
		Interfaceinfoid: interfaceId,
		Leftnum:         leftNum,
	})
}

/** 获得接口调用次数的统计信息，Top N
 */
func ListTopNOfInterfaceInvokeCount(limit int32) ([]*dbsq.ListTopNOfInterfaceInvokeCountRow, error) {
	conn, err := db.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.ListTopNOfInterfaceInvokeCount(ctx, limit)
}

func ConvertSliceToValidTopNOfInterfaceInvokeCountRow(slice []*dbsq.ListTopNOfInterfaceInvokeCountRow) []*models.ValidTopNOfInterfaceInvokeCountRow {
	result := make([]*models.ValidTopNOfInterfaceInvokeCountRow, len(slice))
	for i, item := range slice {
		result[i] = models.Convert2ValidTopNOfInterfaceInvokeCountRow(item)
	}
	return result
}
