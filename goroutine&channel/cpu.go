package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpu=", cpu)

	//可以自己设置使用多个cpu	go1.8之后，默认让程序运行在多核上
	runtime.GOMAXPROCS(cpu-1)
	fmt.Println("ok")
}
