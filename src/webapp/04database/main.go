package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/demo")
	if err != nil {
		fmt.Println("数据库连接错误：", err)
		panic(err.Error())
	}
}

func main() {
	// addRow()
	// selRows(10)
	getRedisRow()
}

func addRow() int64 {
	res, err1 := Db.Exec("insert into a value ('2021-10-07', 10, 10)")
	if err1 != nil {
		fmt.Println("sql执行错误：", err1)
		panic(err1.Error())
	}

	rows, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("受影响行数错误：", err2)
		panic(err2.Error())
	}

	return rows
}

func selRows(a int) {
	rows, err0 := Db.Query("select * from a where a=?", a)
	if err0 != nil {
		fmt.Println(err0)
	}
	for rows.Next() {
		var (
			date string
			a1   int
			b    int
		)
		if err1 := rows.Scan(&date, &a1, &b); err1 != nil {
			fmt.Println(err1)
		}
		fmt.Printf("日期：%v，a=%v, b=%v \n", date, a1, b)
	}
}

func getRedisRow() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.HGet(ctx, "user_data:zms", "112").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)

	// val = `{"server_id":5001,"user_id":112,"level":1,"role_name":"1","pid":"1","gid":"1","create_time":1633591527,"update_time":1633591527,"guild_id":0}`

	type userData struct {
		Level     int
		Server_Id int
	}
	var user1 userData
	json.Unmarshal([]byte(val), &user1)
	fmt.Println("key", user1)

	/*var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	fmt.Println(jsonBlob)
	type Animal struct {
		Name  string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)*/
}
