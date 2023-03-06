package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/faridlan/product-api/helper"
)

func NewDatabase() *sql.DB {

	port := os.Getenv("PORT_DB")
	host := os.Getenv("HOST_DB")
	name := os.Getenv("NAME_DB")

	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s:%s)/%s", host, port, name))
	helper.PanicIfErr(err)

	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
