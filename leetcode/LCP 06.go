package main

import (
	"fmt"
	"time"
)

func main() {
	c := []int{2,3,10}
	st := time.Now()
	str := minCount(c)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func minCount(c []int) int {
	u := 0
	for _,v := range c{
		u += (v+1)/2
	}
	return u
}
