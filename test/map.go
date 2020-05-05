package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {

	a2 := [5]int{0:1, 4:2}
	a21 := [3][2]int{
		[2]int{1,4},
		[2]int{2,5},
		[2]int{3,6},
	}
	fmt.Println(a2,a21)
	a3 := [...]int{1, 3, 5, 7, 8}
	var count1 = 0
	for _,v := range a3 {
		count1 += v
	}
	fmt.Println(count1)

	a4 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a4); i++ {
		numUp := a4[i]
		for u := i+1; u < len(a4); u++ {
			numDown := a4[u]
			sum := numUp + numDown
			if sum == 8 {
				fmt.Println(i, u)
			}
		}
	}

	fmt.Println("----------------------")
	cap1 := []int{1,3,5,7,9,11,13}
	cap2 := cap1[:5]
	cap3 := cap1[3:]
	fmt.Printf("value:%v len(cap2):%d cap(cap2):%d\n", cap2, len(cap2), cap(cap2))
	fmt.Printf("value:%v len(cap3):%d cap(cap3):%d\n", cap3, len(cap3), cap(cap3))
	cap1[4] = 90	//修改底层数组
	fmt.Println(cap2, cap3)
	//切片的本质：
	// 切片就是一个框，框住了一块连续的内存。
	// 切片属于引用类型，真正的数据都是保存在底层数组里的。
	cap4 := make([]int, 5, 10)
	fmt.Printf("value:%v len(cap4):%d cap(cap4):%d\n", cap4, len(cap4), cap(cap4))
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	aCs := []string{"成都", "重庆"}
	citySlice = append(citySlice, aCs...)	// ...表示拆开
	fmt.Printf("value:%v len(citySlice):%d cap(citySlice):%d\n", citySlice, len(citySlice), cap(citySlice)) //[北京 上海 广州 深圳 成都 重庆]

	a:= []int{1,3,5}
	a1 := a
	var a11 = make([]int, 3, 3)
	copy(a11, a)
	fmt.Println(a, a1, a11)

	var arr = [...]int{1,2,3,4,5,4,2,1,5,2,1}
	var a22 = make([]int, len(arr), len(arr))
	for _,v := range arr{
		a22[v]++;
	}
	for k,v := range a22 {
		if v == 1 {
			fmt.Println(k,v)
		}
	}

	s111 := []int{1,2,3,4}
	s112 := s111
	var s113 = make([]int, 1, 3)	//如果长度为0，拷贝元素不会改变长度，所以结果为:[]
	copy(s113, s111)
	fmt.Println(s112)
	s112[1] = 200
	fmt.Println(s112)
	fmt.Println(s111)
	fmt.Println(s113,len(s113), cap(s113))

	// 从切片中删除元素
	a23 := []int{1,3,5,7}	//数组
	s23 := a23[:]			//切片
	// 1、切片不保存具体的值
	// 2、切片对应一个底层数组
	// 3、底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", a23)
	fmt.Printf("%p\n", &s23[0])
	// 要删除索引为2的元素
	s23 = append(s23[:1], s23[2:]...)		//修改了底层数组
	fmt.Printf("%p\n", &s23[0])
	fmt.Println(s23, len(s23), cap(s23))	//从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	fmt.Println(a23, len(a23), cap(a23))

	var a24 = make([]string, 5, 10)		//初始化的时候就有5个空元素
	fmt.Println(a24)	//[         ]
	for i := 0; i < 10; i++ {
		a24 = append(a24, fmt.Sprintf("%v", i))
	}
	fmt.Println(a24)	//[          0 1 2 3 4 5 6 7 8 9]

	var a25 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a25[:])
	fmt.Println(a25)
	fmt.Println("-------------------")

	// 指针
	// Go里面的指针只能读不能修改
	// 1、&:取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)	// *int表示int类型的指针
	// 2、*:根据地址取值
	m := *p
	fmt.Println(m)
	/*
	new与make的区别
	二者都是用来做内存分配的。
	make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。int、string、bool
	*/
	var m1 map[string]int
	fmt.Println(m1==nil)	//还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10)	//要估算好该map容量，避免在程序运行期间再动态扩容
	m1["zhl"] = 20
	m1["age"] = 18
	fmt.Println(m1)
	fmt.Println(m1["zhl"])
	//m1["mumu"] = 0
	if value,ok := m1["mumu"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("没有此key")
	}
	for k,v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}
	for _,v := range m1 {
		fmt.Println(v)
	}
	delete(m1, "age")
	fmt.Println(m1)
	fmt.Println("-----------------")

	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Printf("%v:%v ", key, scoreMap[key])
	}
	fmt.Println("")
	//map和slice组合
	//元素类型为map的切片
	var s1 = make([]map[int]string, 5, 10)	//如果长度为0，调用就会报跨界错误
	s1[0] = make(map[int]string, 1)			//对内部的map做初始化(如果s1 = append(s1, 1)	//自动初始化)
	s1[0][10] = "zhl"
	fmt.Println(s1)
	//值为切片类型的map
	var m11 = make(map[string][]int,10)
	m11["北京"] = []int{10,20,30}			//初始化
	fmt.Println(m11)

	s12:="how do you do"
	// split:=strings.Split(s12," ")
	// update 采用正则匹配来分割字符串
	spaceRe, _ := regexp.Compile(`\s+`)

	split := spaceRe.Split(s12, -1)
	m12:=make(map[string]int,len(split))
	for _,i:= range split{
		value, ok := m12[i]
		if(ok){
			m12[i]=value+1
		}else{
			m12[i]=1
		}
	}
	fmt.Println(m12)

	sq12 := strings.Split(s12, " ")
	var m13 = make(map[string]int, 10)
	for _,v := range sq12 {
		//m13[v]++
		if _,ok := m13[v];ok {
			m13[v]++
		} else {
			m13[v] = 1
		}
	}
	for keys,value := range m13 {
		fmt.Println(keys, value)
	}
}
