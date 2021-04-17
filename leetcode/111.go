package main

import (
	"fmt"
	"time"
)

func main() {
	n := 4421
	st := time.Now()
	str := minDepth(n)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func minDepth(n int) int {
	var j,sum,number int
	j = 1
	for n > 0 {
		number = n%10
		j *= number
		sum += number
		n = n/10
	}
	return j-sum
}
