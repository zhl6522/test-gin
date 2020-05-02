package main

import (
	"fmt"
	"math"
)

const (
	_ = 1*iota
	KB = 1 << (10*iota)
	MB = 1 << (10*iota)
	GB = 1 << (10*iota)
	TB = 1 << (10*iota)
	PB = 1 << (10*iota)
)

func main() {
	fmt.Println(KB, MB)
	var a,b = 3,4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("type:%T,value is:%v\n",c,c)
	var s = "hello沙河小王子"
	temp := []rune(s)
	var count int
	for _,v := range temp {
		if v > 256 {
			count ++
			fmt.Println(string(v))
		}
	}
	fmt.Println(count)
	fmt.Println("-----------------------")
	/*var age = 19
	if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习！")
	}*/
	//作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习！")
	}
	//fmt.Println(age)	//undefined: age
	i := 1
	for i<10 {
		fmt.Print(" ", i)
		i++
	}
	fmt.Println("")
	for a:=1;a<10;a++ {
		for b:=a;b<10;b++ {
			fmt.Printf("%d*%d=%d ", a, b, a*b)
		}
		fmt.Println("\n")
	}
}
