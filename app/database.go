package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/product-api/helper"
)

func NewDatabase() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/nostra")
	helper.PanicIfErr(err)

	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
