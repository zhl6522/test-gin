package groutine

import (
	"runtime"
	"sync"
	"testing"
)

//只需在函数调⽤语句前添加 go 关键字，就可创建并发执⾏单元。开发⼈员⽆需了解任何
//执⾏细节，调度器会⾃动将其安排到合适的系统线程上执⾏。goroutine 是⼀种⾮常轻量
//级的实现，可在单个进程⾥执⾏成千上万的并发任务。
//事实上，⼊⼝函数 main 就以 goroutine 运⾏。另有与之配套的 channel 类型，⽤以实
//现 "以通讯来共享内存" 的 CSP 模式。
//和协程 yield 作⽤类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等
//待下次被调度执⾏。
func TestGroutin(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 { runtime.Gosched() }
		}
	}()
	go func() {
		defer wg.Done()
		println("Hello, World!")
	}()
	wg.Wait()
}
