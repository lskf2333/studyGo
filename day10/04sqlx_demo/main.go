package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/ceshi"
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置数据库连接池的最大空闲连接数
	db.SetConnMaxIdleTime(5)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err:%v", err)
		return
	}
	fmt.Println("连接数据库成功")
	sqlStr1 := `select * from user where id=1`
	var u1 user
	db.Get(&u1, sqlStr1)
	fmt.Printf("u:%v\n", u1)

	var userList []user
	sqlStr2 := `select * from user`
	db.Select(&userList, sqlStr2)
	fmt.Printf("userList:%#v\n", userList)
}
