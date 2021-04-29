package main

import "fmt"

func main() {

	//管道的使用
	//1、创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3) //管道的容量不会自动增加，加多了报错：死锁

	//2、看看intChannel是什么
	fmt.Printf("intChan的值=%v,intChan的地址=%p\n", intChan, &intChan) // 参考同目录文件：9、channel管道介绍6.png

	//3、向管道写入数据
	intChan <- 10 //channel中只能存放指定的数据类型
	num := 211
	intChan <- num

	//4、管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//5、从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//6、在没有使用协程的情况下，如果我们的管道数据已经全部取出，再去就会报告 deadlock

	// 注意：如果向管道里放map数据 需要make
	/*
		var mapChan chan map[string]string
		mapChan = make(chan map[string]string, 10)
		m1 := make(map[string]string, 20)
		m1["city1"] = "北京"
		m1["city2"] = "天津"
		m2 := make(map[string]string, 20)
		m2["hero1"] = "武松"
		m2["hero2"] = "诸葛亮"
		mapChan<-m1
		mapChan<-m2
	*/

}
