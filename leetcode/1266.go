package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	n := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	st := time.Now()
	fmt.Println(minTimeToVisitAllPoints(n))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func minTimeToVisitAllPoints(n [][]int) int {
	m := 0
	for i := 1; i <= len(n); i++ {
		j := i - 1
		max := int(math.Abs(float64(n[i][0] - n[j][0])))
		t := int(math.Abs(float64(n[i][1] - n[j][1])))
		if max < t {
			max = t
		}
		m += max
	}
	return m
}
