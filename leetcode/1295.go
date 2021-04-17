package main

import (
	"fmt"
	"time"
)

func main() {
	n := []int{12, 345, 2, 6, 7896, 78}
	st := time.Now()
	fmt.Println(findNumbers(n))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func findNumbers(n []int) int {
	m := 0
	for _, v := range n {
		num := 0
		for v > 0 {
			num++
			v /= 10
		}
		if num%2 == 0 {
			m++
		}
	}
	return m
}
