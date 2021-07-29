package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/ceshi"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会检验用户名和密码是否正确，只是判断信息格式是否正确
	if err != nil {                   //dsn 格式不正确的时候才会报错
		fmt.Printf("dsn: %s invalid,err:%v", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v", dsn, err)
		return
	}
	fmt.Println("连接数据库成功！")
	fmt.Println(db)
}
