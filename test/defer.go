package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 1、返回值赋值
// 2、真正的RET返回
// 函数如果存在defer，defer执行的时机在1和2之间
func aa(i int) int {
	i+=1
	fmt.Printf("%p\n", &i)
	return i
}
func f1() int {
	x := 5
	//defer aa(x)
	defer func() {
		x++			// 修改的是x不是返回值
		fmt.Printf("x++得到的值：%v，", x)
	}()
	fmt.Printf("%p\n", &x)
	return x		// 1、返回值赋值 2、defer 3、真正的RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5		// 返回值=x=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++			// 修改的是x
	}()
	return x		// 1、返回值 = y = x = 5 2、defer修改的是x 3、真正的返回
}
func f4() (x int) {
	defer func(x int) {
		x++			// 改变的是函数中的副本
	}(x)
	return 5		// 返回值 = x = 5
}
//传一个指针到匿名函数
func f5() (x int) {
	defer func(x *int) {
		(*x)++			// 改变的是函数中的指针
	}(&x)
	return 5		// 返回值 = x = 6
}
func main() {
	fmt.Println(f1())	//5
	fmt.Println(f2())	//6
	fmt.Println(f3())	//5
	fmt.Println(f4())	//5
	fmt.Println(f5())	//6
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	funcA()
	funcB()
	funcC()
}
/*
1、defer clac("AA", 1, calc("A",1, 2))		//A 1 2 3
2、defer clac("AA", 1, 3)
3、defer clac("BB", 10, calc("B", 10,2))	//B 10 2 12
4、defer clac("BB", 10, 12)
5、calc("BB", 10, 12)						//BB 10 12 22
6、clac("AA", 1, 3)							//AA 1 3 4
*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func funcA() {
	fmt.Println("func A")
}
func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
		fmt.Println("释放数据库链接。。。")
	}()
	panic("panic in B")
}
func funcC() {
	fmt.Println("func C")
}