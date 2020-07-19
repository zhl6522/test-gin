package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"		// init()
)
// Go连接MYSQL示例

func main() {
	// 数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/go"
	db, err := sql.Open("mysql", dsn)		// 不校验用户名和密码是否正确
	if err != nil {
	    fmt.Printf("dsn:%s failed, err:%v\n",dsn, err)
	    return
	}
	err = db.Ping()		// 尝试连接数据库
	if err != nil {
	    fmt.Printf("open %s failed, err:%v\n",dsn, err)
	    return
	}

	fmt.Println("success")
}
