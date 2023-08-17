package service

import (
	"context"
	"database/sql"
	"xj/xapi-backend/db"
	"xj/xapi-backend/dbsq"
	"xj/xapi-backend/models"
)

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

func CreateInterface(param *models.CreateInterfaceParams) (sql.Result, error) {
	nparam := models.ConvertToCreateInterfaceParams(param)
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.CreateInterface(ctx, nparam)
}

func UpdateInterface(param *models.UpdateInterfaceParams) error {
	_, err := GetInterfaceInfo(param.ID)
	if err != nil {
		return err
	}
	nparam := models.ConvertToUpdateInterfaceParams(param)
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	return q.UpdateInterface(ctx, nparam)
}

func DeleteInterface(id int64) error {
	q := dbsq.New(db.MyDB)
	ctx := context.Background()
	_, err := GetInterfaceInfo(id)
	if err != nil {
		return err
	}
	return q.DeleteInterface(ctx, id)
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
