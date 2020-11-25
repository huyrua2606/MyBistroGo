package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func connectDB() {
	db, err = sql.Open("mysql", "root:Quang123Huy@@/mybistro")
	fmt.Println("Database connected.")
	if err != nil {
		panic(err.Error())
	}
}
