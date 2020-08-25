package main

import (
	"fmt"
	"time"
)

/*
输入：guess = [1,2,3], answer = [1,2,3]
输出：3
解释：小A 每次都猜对了。

输入：guess = [2,2,3], answer = [3,2,1]
输出：1
解释：小A 只猜对了第二次。
*/
func main() {
	st := time.Now()
	guess := []int{2, 2, 3}
	answer := []int{3, 2, 1}
	ret := game(answer, guess)
	t := time.Since(st)
	fmt.Println(ret, t)
}

func game(answer, guess []int) int {
	i := 0
	for k, v := range answer {
		if v == guess[k] {
			i++
		}
	}
	return i
}
