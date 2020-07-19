package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql" // init()
)
// Go连接MYSQL示例

var db *sqlx.DB		// 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/go"
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(5)	// 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(3)	// 设置连接池中的最大闲置连接数
	return
}

type user struct {
	Id		int
	Name 	string
	Age 	int
}

// sql注入示例
func sqlInjectDemo(name string) {
	// 自己拼接sql语句的字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", users)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n",err)
		return
	}
	fmt.Println("连接数据库成功！")
	// SQL注入几种示例
	//sqlInjectDemo("黎明")
	//sqlInjectDemo("XXX' or 1=1 #")
	sqlInjectDemo("XXX' union select * from user #")
}
