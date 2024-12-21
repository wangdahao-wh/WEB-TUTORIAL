package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dns := "root:wh985211@tcp(127.0.0.1:3306)/itheima"
	var err error
	db, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success")
	id := 2
	account, err := QueryRowDemo(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
