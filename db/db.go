package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

const (
	MaxConns = 100
	MinConns = 2
)

func init() {
	db, err = sql.Open("mysql", "root:mysql3306recycle072829@(112.124.31.82:3306)/poem?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(MaxConns)
	db.SetMaxOpenConns(MinConns)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
