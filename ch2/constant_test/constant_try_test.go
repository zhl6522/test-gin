package main

import (
	//"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a:=3	//0011
	t.Log(Readable, Writable, Executable)
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}