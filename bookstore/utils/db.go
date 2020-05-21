package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:123456@/godb")
	if err != nil {
		panic(err)
	}

}
