package mariaDB

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

var db, nil = sql.Open("mysql", "root@/atonproject?charset=utf8")

func Select(query string) *sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func Insert(query string) sql.Result {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	return res
}