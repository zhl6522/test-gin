package main

import "fmt"

func f1(a ...interface{}) {
	fmt.Printf("type:%T value:%#v\n", a, a)
}

func main() {
	//f1()		// type:[]interface {} value:[] 空接口类型的切片
	//f1(1)
	//f1(1,false, "a", struct {}{}, []int{1,2}, [...]int{1,2,3}, map[string]int{"zhl":25})

	var s = []interface{}{1,2,3,4}
	f1(s)
	f1(s...)
}
