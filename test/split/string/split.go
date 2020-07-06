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
	var ret []string
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
