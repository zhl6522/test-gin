package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁
// rwlock
var x = 0
//var wg sync.WaitGroup
//var lock sync.Mutex		// 互斥锁

var (
	y = 0
	wg		sync.WaitGroup
	lock	sync.Mutex		// 互斥锁
	rwlock	sync.RWMutex	// 读写互斥锁
	// 如果读远远大于写的时候，读写互斥锁比互斥锁效率高
)

//var smap sync.Map			// Go语言中内置的map不是并发安全的。

// 读操作
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(y)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	wg.Done()
	rwlock.Lock()
	y++
	time.Sleep(time.Millisecond*5)
	rwlock.Unlock()
}

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()		// 加锁
		x = x + 1
		lock.Unlock()	// 解锁
	}
	wg.Done()
}

func main() {
	/*wg.Add(2)
	go add()		// 可能两者同时读到50，也同时返回50
	go add()
	wg.Wait()
	fmt.Println(x)*/
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
