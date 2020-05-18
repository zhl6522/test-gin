package mylogger

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	pc,file,line,ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName,file,path.Base(file),line)
}

func f1() {
	f()
	_,file,line,ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
	}
	fmt.Println(file,line)
}

func main() {
	f1()
	_,file,line,ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed")
	}
	fmt.Println(file,line)
}
