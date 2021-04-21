package main

import (
	"encoding/json"
	"fmt"
)

// 类型转换

func main() {
	// "hello@guigu"转成"好ello@guigu"
	str := "hello@guigu"
	fmt.Println("str=", str)
	arr1 := []rune(str)		// []byte()只能处理单字节的数字和字母，一个汉字3个字节处理会乱码
	arr1[0] = '好'
	str = string(arr1)
	fmt.Println("str=", str)

	//字符串和数组的转换
	arr := []rune(str)
	fmt.Printf("字符串转数组：类型：%T 值：%c \n", &arr, arr)		//然后进行 for range操作
	var data [10]byte
	data[0] = 'T'
	fmt.Printf("数组转字符串：%s \n", string(data[:]))
	fmt.Printf("数组转字符串：%s \n", string(arr))

	arr3 := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, _ := json.Marshal(arr3)
	fmt.Printf("array 到 json str：%s \n", lang)


	//切片和数组的转换
	s := []int{1,2,3}
	var arr2 [3]int
	copy(arr2[:], s)
	fmt.Printf("切片转数组：类型：%T 值：%v\n", &arr2, arr2)
	s2 := make([]int, 3)
	copy(s2, arr2[:])
	fmt.Printf("数组转切片：类型：%T 值：%v\n", &s, s)
}
