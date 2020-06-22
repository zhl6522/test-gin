package main

import (
	"fmt"
	//"strconv"
	"sync"
	"sync/atomic"
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
	lock	sync.Mutex		// 互斥锁		防止同一时刻多个goroutine操作同一个资源。
	rwlock	sync.RWMutex	// 读写互斥锁	适用于度多写少的场景，才能提高程序的执行效率。	特点：1、读的goroutine来了获取的是读锁，后续的goroutine能读不能写。2、写的goroutine来了获取的是写锁，后续的goroutine不管是读是写都要等待获取锁。
	// 如果读远远大于写的时候，读写互斥锁比互斥锁效率高
)

var m2 = sync.Map{}			// 开箱即用（不需要make初始化）的并发安全的Map

//var smap sync.Map			// Go语言中内置的map不是并发安全的。 使用并发访问的Map；同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
//var once sync.Once		//	案例：test/channel/main.go 某些函数只需要执行一次的时候，就可以使用sync.Once
							// 比如并发

var z	int64

// 读操作
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(y)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
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

func add2() {			// 原子操作 Go语言内置了一些针对内置的基本数据类型的一些并发安全的操作
	atomic.AddInt64(&z, 1)
	wg.Done()
}

func main() {
	/*wg.Add(2)
	go add()		// 可能两者同时读到50，也同时返回50
	go add()
	wg.Wait()
	fmt.Println(x)*/

	/*start := time.Now()
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
	fmt.Println(time.Now().Sub(start))*/

	/*wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)		// 必须使用sync.Map内置的Store方法设置键值对
			value,_ := m2.Load(key)	// 必须使用sunc.Map提供的Load方法根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()*/

	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add2()
	}
	wg.Wait()
	fmt.Println(z)

	// 比较并交换
	z = 100
	ok := atomic.CompareAndSwapInt64(&z, 100, 200)
	fmt.Println(ok, z)
}
