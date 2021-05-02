package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for v := 1; v < 80000; v++ {
		prime := 1
		put := v / 2
		for i := 2; i <= put; i++ {
			if v%i == 0 {
				prime = 0
				break
			}
		}
		if prime == 1 {

		}
	}
	end := time.Now().UnixNano()
	fmt.Println("不使用协程的普通方法耗时=", end-start)
}
