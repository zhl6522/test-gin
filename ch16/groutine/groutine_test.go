package groutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// M -- System Thread 「系统线程」
// P -- Processor 「Go 语言的协程处理器」
// G -- Goroutine 「协程」
func TestGroutine(t *testing.T) {
	for i:=0;i<10;i++ {
		go func(par int) {	// 使用go创建协程,但是需要注意的是：协程函数的par作为参数是外部i的数据拷贝
			fmt.Println(par)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func doSomething(i int) {
	fmt.Print(i)
}

func TestSchedtick(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			doSomething(0)
		}
	}()
	for {
		doSomething(1)
	}
}