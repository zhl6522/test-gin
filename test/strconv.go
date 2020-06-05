package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 从字符串中解析出相应的数据
	str := "1000"
	//ret1 := int64(str)
	ret1,err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("ParseInt failed,err:%v\n", err)
		return
	}
	fmt.Printf("%#v %T\n",ret1, int(ret1))

	// 字符串转换成整形
	retInt,_ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt,retInt)

	// 从字符串中解析布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

	// 把数字转换成字符串类型
	i := int(97)
	//ret1 := string(i)	//"97"
	//fmt.Println(ret1)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v %T\n",ret2, ret2)

	// 整形转换成字符串
	ret3 := strconv.Itoa(retInt)
	fmt.Printf("%#v %T\n", ret3,ret3)

}
