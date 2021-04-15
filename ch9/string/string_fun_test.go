package string__test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {
	s := "A:B:C"
	split := strings.Split(s,":")
	t.Log(split)
	for _,part := range split {
		t.Log(part)
	}
	t.Log(strings.Join(split, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)	//strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。
	t.Log("str" + s)
	if i,err := strconv.Atoi("10");err == nil {	//strconv.Atoi()函数的参数是一个字符串类型的数字，它可以将字符串类型的数字转换成对应的整型数字。
		t.Log(10 + i)
	}
}