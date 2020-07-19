package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sql.DB		// 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/go"
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)		// 不校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping()		// 尝试连接数据库
	if err != nil {
		return
	}
	return
}

func transactionDemo() {
	// 1、开启事务
	tx, err := db.Begin()
	if err != nil {
	    fmt.Printf("begin failed, err:%v\n",err)
	    return
	}
	// 2、执行多个sql操作
	sqlstr1 := `update user set age=age-2 where id=1`
	sqlstr2 := `update xxx set age=age+2 where id=5`
	_, err = tx.Exec(sqlstr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Printf("update failed, err:%v\n",err)
		return
	}
	_, err =tx.Exec(sqlstr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Printf("update failed, err:%v\n",err)
		return
	}
	// 上面两部都执行成功， 提交本次事务
	err = tx.Commit()
	if err != nil {
	    fmt.Println("提交出错啦 要回滚")
	    return
	}
	fmt.Println("事务执行成功！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n",err)
		return
	}
	fmt.Println("连接数据库成功！")
	transactionDemo()
}
