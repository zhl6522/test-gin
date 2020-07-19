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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n",err)
		return
	}
	fmt.Println("连接数据库成功！")
	sqlStr1 := `select id,name,age from user where id=1`
	var u user
	db.Get(&u, sqlStr1)
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name,age from user`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
	    fmt.Printf("select failed, err:%v\n",err)
	    return
	}

	fmt.Printf("userList:%#v\n", userList)
}
