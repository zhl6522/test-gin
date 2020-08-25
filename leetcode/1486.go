package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ret := xor(5, 10)
	t := time.Since(st)
	fmt.Println(ret, t)
}

func xor(start, n int) int {
	nums := start
	for i := 1; i < n; i++ {
		num := start + i*2
		nums = nums ^ num
	}
	return nums
}
