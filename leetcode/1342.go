package main

import (
	"fmt"
	"time"
)

func main() {
	n := 59
	st := time.Now()
	str := numberOfSteps(n)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func numberOfSteps(s int) int {
	n := 0
	for s != 0 {
		if s%2 == 0 {
			s /= 2
		} else {
			s -= 1
		}
		n++
	}
	return n
}
