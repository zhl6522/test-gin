package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	n := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	st := time.Now()
	fmt.Println(getDecimalValue(n))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func getDecimalValue(n []int) int {
	m := 0
	lenn := len(n)
	for i := 0; i < lenn; i++ {
		if n[i] == 1 {
			le := lenn - i - 1
			pow := math.Pow(float64(2), float64(le))
			m += int(pow)
		}
	}
	return m
}
