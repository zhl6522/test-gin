package main

import (
	"fmt"
	"time"
)

func main() {
	n := []int{8, 1, 2, 2, 3}
	st := time.Now()
	fmt.Println(smallerNumbersThanCurrent(n))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func smallerNumbersThanCurrent(n []int) []int {
	num := make([]int, len(n))
	for k, _ := range n {
		for k2, _ := range n {
			if n[k] > n[k2] {
				num[k]++
			}
		}
	}
	return num
}
