package db

import (
	"database/sql"
	"fmt"
)

var MyDB *sql.DB

func ConnectionPool(savePath string) *sql.DB {
	db, err := sql.Open("mysql", savePath)
	if err != nil {
		fmt.Println("sql.Open err", err)
		return nil
	}
	return db
}
