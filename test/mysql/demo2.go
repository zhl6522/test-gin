package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)
// Go连接MYSQL示例

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
	db.SetMaxOpenConns(5)	// 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(3)	// 设置连接池中的最大闲置连接数
	return
}

type user struct {
	id		int
	name 	string
	age 	int
}

// 查询
func queryRow(id int) {
	var u1 user
	// 1、写查询条件的记录的sql语句
	sqlStr := `select id,name,age from user where id=?`
	// 2、执行并拿到结果
	/*for i := 0; i < 7; i++ {
		db.QueryRow(sqlStr, 3)		// 从连接池里拿一个连接出来去数据库查询记录
		fmt.Printf("已经执行第%d次查询\n", i)
	}*/
	// 从连接池里拿一个连接出来去数据库查询记录		// 必须对rowObj对象调用Scan方法，因为该方法会释放数据库连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	// 4、打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条数据示例
func queryMultiRowDemo(n int)  {
	var u1 user
	// 1、SQL语句
	sqlStr := `select id,name,age from user where id>?`
	// 2、执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
	    fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
	    return
	}
	// 3、一定要关闭rows
	defer rows.Close()
	// 4、循环取值
	for rows.Next() {
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
		    fmt.Printf("scan failed, err:%v\n",err)
		    return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insertDB() {
	sqlStr := `insert into user(name,age) values('晶晶', 27)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
	    fmt.Printf("insert failed, err:%v\n",err)
	    return
	}
	// 如果是插入数据的操作，能拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
	    fmt.Printf("get last id failed, err：%v\n",err)
	    return
	}
	fmt.Println("id:", id)
}

// 更新操作
func updateRow(id, age int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n",err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update failed, err：%v\n",err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除
func deleteRow(id int) {

	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n",err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete id failed, err：%v\n",err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 余出来插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
	    fmt.Printf("prepare failed, err：%v\n",err)
	    return
	}
	defer stmt.Close()
	// 后续只需要拿到stmt执行一些操作
	var m = map[string]int{
		"张三":21,
		"李四":22,
		"王五":23,
	}
	for k,v := range m{
		_, err :=stmt.Exec(k, v)		// 后续只需要处理
		if err != nil {
		    fmt.Printf("prepare failed, err:%v\n",err)
		    return
		}
	}
}

func main() {
	err := initDB()
	if err != nil {
	    fmt.Printf("init DB failed, err:%v\n",err)
	    return
	}
	fmt.Println("连接数据库成功！")
	queryRow(1)
	queryMultiRowDemo(2)
	//insertDB()
	updateRow(1, 30)
	//deleteRow(5)
	prepareInsert()
}
