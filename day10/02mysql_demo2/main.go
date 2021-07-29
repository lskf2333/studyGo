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

//查询单个记录
func queryOne(id int) {
	var u1 user
	//1. 写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id =?`
	// // 2.执行
	// rowObj := db.QueryRow(sqlStr, 2) //从连接池中拿一个连接出来去数据库查询单条记录
	// // 3.拿到结果
	// rowObj.Scan(&u1.id, &u1.name, &u1.age)
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	//打印结果
	fmt.Printf("u1: %#v", u1)
}

//查询多条记录
func queryMore(n int) {
	// 1. sql 语句
	sqlStr := `select * from user where id > ?`
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("xexc %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 3.一定要关闭rows
	defer rows.Close()
	// 4.循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err%v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}

}

func insert() {
	// 1. 写sql 语句
	sqlStr := `insert into user(name,age) value ("lala",20)`
	// 2. 执行
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//如果是插入数据的操作，能拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

//更新
func updateRow(newAge int, id int) {
	// 1. 写sql 语句
	sqlStr := `update user set age=? where id=?`
	// 2. 执行
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

//删除
func deleteRow(id int) {
	// 1. sql
	sqlStr := `delete from user where	id =?`
	// 2. 执行
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

//预处理
func prepareInsert() {
	sqlStr := `insert into user (name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"kaka": 21,
		"gaga": 22,
		"fafa": 23,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err:%v", err)
		return
	}
	fmt.Println("连接数据库成功")
	// queryOne(2)
	// insert()
	// updateRow(90, 2)
	// deleteRow(3)
	prepareInsert()
	queryMore(0)
}
