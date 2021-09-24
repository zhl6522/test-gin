package main

import (
	"fmt"
	"unsafe"
)

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


	//https://www.liwenzhou.com/posts/Go/struct_memory_layout/
	//总结一下：在了解了Go的内存对齐规则之后，我们在日常的编码过程中，完全可以通过合理地调整结构体的字段顺序，从而优化结构体的大小。
	var f Foo
	fmt.Println(unsafe.Sizeof(f))  // 3

	var b1 Bar
	fmt.Println(unsafe.Sizeof(b1)) // 24

	var b2 Bar2
	fmt.Println(unsafe.Sizeof(b2)) // 16
}

type Foo struct {
	A string
}

type Bar struct {
	x int32 // 4
	y *Foo  // 8
	z bool  // 1
}

type Bar2 struct {
	x int32 // 4
	z bool  // 1
	y *Foo  // 8
}