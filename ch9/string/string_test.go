package string__test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)	//初始化为默认零值“”
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' //string是不可变的byte slice，不能赋值
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))	//是byte数

	c:=[]rune(s)	//rune代表着Unicode
	t.Log(len(c))	//此处不是byte数
	//t.Log("run size:", unsafe.Sizeof(c[0]))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _,v := range s {
		t.Logf("%[1]c %[1]x", v)
	}
}
