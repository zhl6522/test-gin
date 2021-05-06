package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	vVal := reflect.ValueOf(v)
	fmt.Printf("vVal.Value=%v\n", vVal)
	fmt.Printf("vVal.Type=%v\tvVal.Kind=%v\n", reflect.TypeOf(v), vVal.Kind())
	vI := vVal.Interface()
	fmt.Printf("vI's float64=%v\n", vI.(float64))

	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("str改变后的值=%v\n", str)
}
