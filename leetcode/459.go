package main

import (
	"fmt"
	"time"
)

func main() {
	n := "abcabcabcabc"
	st := time.Now()
	str := repeatedSubstringPattern(n)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break;
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
