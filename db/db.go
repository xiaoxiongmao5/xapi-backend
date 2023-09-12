package db

import (
	"database/sql"
)

var MyDB *sql.DB

func ConnectionPool(savePath string) (*sql.DB, error) {
	return sql.Open("mysql", savePath)
}
