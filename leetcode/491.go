package main

import (
	"fmt"
	"time"
)

func main() {
	//n := []int{4, 6, 7, 7, 8}
	st := time.Now()
	fmt.Println(1<<3)
	//fmt.Println(findSubsequences(n))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func findSubsequences(n []int) [][]int {
	m := make([][]int, 0)
	for i := 0; i < len(n); i++ {
		m = sub(n[i+1:])
		fmt.Println(m)
	}
	return m
}

func sub(n []int) [][]int {
	m := make([][]int, 0)
	for i := 0; i < len(n); i++ {

	}
	return m
}
