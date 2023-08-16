package service

import (
	"context"
	"database/sql"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/models"
)

func CreateInterface(param *dbsq.CreateInterfaceParams) (sql.Result, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.CreateInterface(ctx, param)
}

func ConvertSliceToValidXapiInterfaceInfo(slice []*dbsq.XapiInterfaceInfo) []*models.ValidXapiInterfaceInfo {
	result := make([]*models.ValidXapiInterfaceInfo, len(slice))
	for i, item := range slice {
		result[i] = models.ConvertToValidXapiInterfaceInfo(item)
	}
	return result
}

func ListInterfaces() ([]*models.ValidXapiInterfaceInfo, error) {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	res, error := q.ListInterfaces(ctx)
	if error != nil {
		return nil, error
	}
	return ConvertSliceToValidXapiInterfaceInfo(res), nil
}
