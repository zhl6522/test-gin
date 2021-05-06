package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Score   float32
	Address string
}
//方法，显示s的值
func (m Monster) Print() {
	fmt.Println("-------------")
	fmt.Println(m)
	fmt.Println("-------------")
}
//方法，返回两个数的和
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
//方法，接收4个值，给m赋值
func (m Monster) Set(name string, age int, score float32, address string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Address = address
}
func TestStruct(a interface{}) {
	//获取reflect.type类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	fmt.Printf("kd type=%v\n",kd)
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)	//4

	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d 值为%v\n", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d tag为%v\n", i, tagVal)
		}
	}

	//获取该结构体有几个方法
	getNumMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", getNumMethod)

	//方法的排序默认是按照 函数名的排序(ASCII码)
	val.Method(1).Call(nil)	//获取到第二个方法并调用它
	//调用结构体的第一个方法
	var params []reflect.Value	//声明 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)	//传入的参数是 []reflect.Value 返回 []reflect.Value
	fmt.Printf("res=%v\n", res[0].Int())	//返回结果是 []reflect.Value，如果不确定res[0]的类型，使用断言
}

func main()  {
	var monster = Monster{
		Name:"小狐狸",
		Age:500,
		Score:99.9,
	}
	TestStruct(monster)
}
