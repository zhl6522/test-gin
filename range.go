package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [3]int{0, 1, 2}
	fmt.Printf("%T", a)		//数组Array
	fmt.Println(reflect.TypeOf(a))	//数组:值类型。int,float,bool,string,以及数组和struct，特点：变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 // 使⽤复制品中取出的 value 修改原数组。
	}
	fmt.Println(a) // 输出 [100, 101, 102]。


	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T", s)		//切片Slice
	fmt.Println(reflect.TypeOf(s))	//切片:引用类型。指针，slice，map，chan等都是引用类型，特点：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，通过GC回收。
	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s = s[:3] // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}

	ss := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := ss[2:6:7] // [2 3 4 5]
	fmt.Println(s1, len(s1), cap(s1))	// [2 3 4 5] 4 5
	s2 := ss[2:6] // [2 3 4 5]
	fmt.Println(s2, len(s2), cap(s2))	// [2 3 4 5] 4 8

	s0 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s10 := s0[2:5] // [2 3 4]
	s10[2] = 100
	fmt.Println(s10, len(s10), cap(s10))	// [2 3 100] 3 8
	s20 := s10[2:6] // [100 5 6 7]
	s20[3] = 200
	fmt.Println(s0)
}
