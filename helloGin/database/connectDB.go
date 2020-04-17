package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDataBase() *sql.DB {
	url := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
