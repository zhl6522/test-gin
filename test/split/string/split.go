package string

import (
	"fmt"
	"strings"
)

// Split 切割字符串
// example:
// abc, b => [a,c]
func Split(str string, sep string) []string {
	// "babcdef"	sep:"b"		[a cdef]
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	if index == -5 {
		fmt.Println("no")
	}
	ret = append(ret, str)
	return ret
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}