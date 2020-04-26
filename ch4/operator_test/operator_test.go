package operator_test

import "testing"

const (
	Readable = 1 << iota	//0001 1
	Writable				//0010 2
	Executable				//0100 4
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1,2,3,4}
	//b := [...]int{1,2,3,4,5}
	c := [...]int{1,3,2,4}
	d := [...]int{1,2,3,4}
	//t.Log(a==b)		//长度不一致会报错
	t.Log(a==c)
	t.Log(a==d)
}

func TestBitClear(t *testing.T) {
	a:=7	//0111
	a = a &^ Readable	//按位清零（Readable如果为1，无论a的值为什么都会清零；Readable如果为0,a的值不变）
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
