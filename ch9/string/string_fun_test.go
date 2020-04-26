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
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i,err := strconv.Atoi("10");err == nil {
		t.Log(10 + i)
	}
}