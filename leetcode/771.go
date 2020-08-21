package main

import (
	"fmt"
	"time"
)

func main() {
	j := "aAz"
	s := "aAAbbrgdwet4gqvrsdfdsbdzaqbb"
	st := time.Now()
	len := numJewelsInStones(j,s)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(len)
}

func numJewelsInStones(j,s string) int {
	u := 0
	for _,v1 := range j{
		for _,v2 := range s{
			fmt.Println(v1,v2)
			if v1 == v2 {
				u++
			}
		}
	}
	return u
}
