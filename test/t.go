package main

import (
	"fmt"
	"math"
	"unicode"
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
	var ct int
	for _,c := range s {
		if unicode.Is(unicode.Han, c) {
			ct++
		}
	}
	fmt.Println(ct)
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
		fmt.Println(" ")
	}
	fmt.Println("-------------------------")
	switch d := 7; d {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
		fallthrough		//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。了解
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("未知")
	}

	var boo = false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				boo = true
				break
				// 设置退出标签
				//goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		if boo {
			break
		}
	}
	// 标签
/*breakTag:
	fmt.Println("结束for循环")*/

	// 回文判断：上海自来水来自海上、黄山落叶松叶落山黄
	ss := "a上海自来水来自海上a"	//一个汉子三个字符
	r := make([]rune, 0, len(ss))	//切片
	for _,v := range ss {
		r = append(r, v)
	}
	//fmt.Println(r)
	for i:=0;i<len(r);i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
		}
	}
	fmt.Println("回文")
}
