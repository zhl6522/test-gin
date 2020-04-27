package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if v,ok := p.(string); ok {
	//	fmt.Println("String", v)
	//	return
	//}
	//fmt.Println("Unkonw Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(nil)
}