package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改 num int的值，修改student的值
func reflect01(b interface{}) {
	//1、获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\trVal=%v\trVal's type=%T\n", rVal.Kind(), rVal, rVal) //rVal.Kind()是一个指针
	//Elem返回v持有的接口保管的值的Value的封装，或者v持有的指针指向的Value封装
	rVal.Elem().SetInt(101)	//SetInt()的v必须是一个value

	//这里理解 rVal.Elem()
	num :=9
	var pt *int= &num
	*pt=3		//类似 rVal.Elem()
	//num2 := *pt
	fmt.Println("num2=", *pt)

}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("改变后的num=", num)
}
