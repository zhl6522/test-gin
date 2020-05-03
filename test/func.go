package main

import "fmt"

func sum(x int,y int)(ret int) {
	ret = x + y
	return
	//return x + y
}
func f5() (int, string) {
	return 1,"zhl"
}
func f6(x, y int,m,n string,i,j bool) int {		//参数类型简写：连续多个参数的类型一致时，可以将非最后一个参数的类型省略
	return x + y
}
func f7(x string, y ...int) {	//可变长参数
	fmt.Println(x)
	fmt.Println(y)		// y的类型是切片 []int
}
func main() {
	a := sum(1,2)
	fmt.Println(a)
	m, n:=f5()
	//_, n:=f5()
	fmt.Println(m,n)
	f7("zhl", 1,3,5,7)

}
