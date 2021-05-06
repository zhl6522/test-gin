package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflaceTest01(b interface{}) {
	//通过反射获取传入的变量的type,kind,值
	//1、先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)
	//2、获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//n2 := 2 + rVal	//rVal的类型不是int，会报错(invalid operation: 2 + rVal (mismatched types int and reflect.Value))
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v\trVal's type=%T\n", rVal, rVal)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV=%v\tiV's type=%T\n", iV, iV)
	//将 interfeace{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v\tnum2's type=%T\n", num2, num2)
}

type Student struct {
	Name string
	Age  int
}

//专门演示反射(对结构体的反射)
func reflaceTest02(b interface{}) {
	//通过反射获取传入的变量的type,kind,值
	//1、先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)
	//2、获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//3、获取变量对应的kind
	//(1)、rVal.Kind() ==>
	//(2)、rTyp.Kind() ==>
	fmt.Printf("rVal.Kind=%v rTyp.Kind=%v\n",rVal.Kind(), rTyp.Kind())	//类别：struct

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()		//iV的类型：main.Student
	fmt.Printf("iV=%v\tiV's type=%T\n", iV, iV) //编译阶段 是没办法知道这个类型是什么，所以这里调用iV.Name会报错
	//将 interfeace{} 通过断言转成需要的类型
	/*
	//如果不知道类型 可以使用switch
	switch iV.(type) {
	case bool:
		fmt.Println("bool")
	case int,int64:
		fmt.Println("int")
	default:
		fmt.Println("unknow")
	}*/
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu=%v\tstu's type=%T stu.Name=%v\n", stu, stu, stu.Name) //编译阶段 是没办法知道这个类型是什么，所以这里调用iV.Name会报错
	}
}

func main() {
	//编写一个案例，演示对（基本数据类型，interface{}、reflect.Value）进行反射的基本操作
	//1、先定义一个int
	var num int = 100
	reflaceTest01(num)
	fmt.Println("------------------Student----------------")
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflaceTest02(stu)
}
