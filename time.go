package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	times := fmt.Sprintf("%d-%d-%d %d:%d:%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(), now.Second())
	fmt.Println(times)
	time2 := now.Format("2006-01-02 15:04:05")
	fmt.Println(time2)

	start := now.Unix()
	test03()
	end := now.Unix()
	fmt.Printf("执行test03耗费%v秒\n", end-start)
}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello"+strconv.Itoa(i)
	}
}