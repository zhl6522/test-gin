package main

import (
	"fmt"
	"sort"
)

func main() {
	//map默认是无序的
	map1 := make([]map[string]string, 2) //声明map切片
	if map1[0] == nil {
		map1[0] = make(map[string]string, 2)
		map1[0]["name"] = "牛魔王"
		map1[0]["age"] = "500"
	}
	if map1[1] == nil {
		map1[1] = make(map[string]string, 2) //这一步不能省略
		map1[1]["name"] = "玉兔精"
		map1[1]["age"] = "400"
	}
	// map的容量达到后，不能再像上面这样增加了
	/*if map1[2] == nil {
		map1[2] = make(map[string]string, 2) //这一步不能省略
		map1[2]["name"] = "玉兔精"
		map1[2]["age"] = "400"
	}
	//panic: runtime error: index out of range [2] with length 2
	*/
	mapNew := map[string]string{
		"name": "铁扇公主",
		"age":  "350",
	}
	map1 = append(map1, mapNew)
	fmt.Println(map1)

	//map排序
	map2 := make(map[int]int, 10)
	map2[10] = 100
	map2[4] = 60
	map2[9] = 30
	map2[1] = 10
	fmt.Println(map2)
	
	/*
		如何按照map的key的顺序进行顺序输出
		1、先将map的key放到切片中
		2、对切片排序
		3、遍历切片，然后按照key来输出map的值
	*/
	var keys []int
	for k, _ := range map2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("map2[%v]=%v\n", k, map2[k])
	}

	//map是引用类型，准守引用类型传递的机制，在一个函数接收map,
	//修改后，会直接修改原来的map
	modify(map2)
	fmt.Printf("修改后的map2:%v\n", map2)
}

func modify(map2 map[int]int) {
	map2[9] = 88
}
