package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

//递归应用场景：递归分析1/2/3/4.png
func main() {
	n := 4
	test(n)
}
