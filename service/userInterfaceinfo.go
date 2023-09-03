package service

import (
	"context"
	"database/sql"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
)

/** 用户调用接口关系-计数变更
 */
func InvokeCount(param *dbsq.InvokeUserInterfaceInfoParams) (sql.Result, error) {
	// // 判断用户是否存在
	// // 判断接口是否存在
	// if interfaceId <= 0 || userId <= 0 {
	// 	return nil, errors.New("用户或接口不存在")
	// }
	// 计数变更
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.InvokeUserInterfaceInfo(ctx, param)
}
