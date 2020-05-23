package main

import (
	"fmt"
	"reflect"
)

type MysqlConfig struct {
	Address		string	`int:"address"`
	Port		int		`int:"port"`
	Username	string	`int:"username"`
	Password	string	`int:"password"`
}

func loadIni(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {

	}
	// 反射中使用 Elem()方法获取指针对应的值
	/*if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}*/
}

func main() {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Println(mc.Address,mc.Port,mc.Username,mc.Password)
}
