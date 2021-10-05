package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/demo")
	if err != nil {
		fmt.Println("错误：", err)
		panic(err.Error())
	}
	fmt.Println(Db)
}

func main() {

}
