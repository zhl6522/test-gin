package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ret := leftwords("abcdefghi", 3)
	t := time.Since(st)
	fmt.Println(ret, t)
}

func leftwords(s string, n int) string {
	s1 := s[:n]
	s2 := s[n:]
	return s2+s1
}
