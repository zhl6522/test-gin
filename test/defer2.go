package main

import "fmt"

//当defer被声明时，其后面导函数参数会被实时解析
func main() {
	var i int=1
	//fmt.Println在defer后面，它的参数会实时计算
	// 输出：result=>2(而不是4)
	defer fmt.Println("result1=>", func() int {
		fmt.Println("result1 i=>", i)
		return i*2
	}())
	i++

	//下面的defer后面的函数无参数，所以最里层的i应该是4
	defer func() {
		fmt.Println("result2 i=>", i)
		fmt.Println("result2=>", i*2)
	}()
	i++

	defer fmt.Println("result3=>", func() int {
		fmt.Println("result3 i=>", i)
		return i*2
	}())
	i++
}