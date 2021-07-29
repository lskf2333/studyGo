package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/ceshi"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会检验用户名和密码是否正确，只是判断信息格式是否正确
	if err != nil {                  //dsn 格式不正确的时候才会报错
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置数据库连接池的最大空闲连接数
	db.SetConnMaxIdleTime(5)
	return
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v", err)
		return
	}
	//执行多个sql操作
	sqlStr1 := `update user set age =age-2 where id=1`
	sqlStr2 := `update user set age =age+2 where id=2`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback() //事务回滚
		fmt.Println("执行sql1出错了，要回滚！")
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback() //事务回滚
		fmt.Println("执行sql2出错了，要回滚！")
	}
	//上面两步都执行成功了，就提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错了，要回滚！")
		return
	}
	fmt.Println("事务执行成功了")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err:%v", err)
		return
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}
