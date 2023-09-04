package service

import (
	"context"
	"database/sql"
	"errors"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
)

/** 用户调用接口关系-计数变更
 */
func InvokeCount(interfaceId, userId int64) (sql.Result, error) {
	// 判断接口和用户是否存在
	if interfaceId <= 0 || userId <= 0 {
		return nil, errors.New("用户或接口不存在")
	}
	// 计数变更
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	param := &dbsq.InvokeUserInterfaceInfoParams{
		Interfaceinfoid: interfaceId,
		Userid:          userId,
	}
	return q.InvokeUserInterfaceInfo(ctx, param)
}
