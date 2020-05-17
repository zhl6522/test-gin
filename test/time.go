package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()		// 获取本地的时间
	fmt.Println(now,now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(now.Date())
	// 时间戳 UnixNano纳秒数
	fmt.Println(now.Unix(),now.UnixNano())
	// time.Unix()
	ret := time.Unix(1589706521, 0)
	fmt.Println(ret, ret.Year(),ret.Month(),ret.Day(),ret.Hour(),ret.Minute(),ret.Second())
	fmt.Println(ret.Date())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(time.Hour))
	// Sub 两个时间想减
	// 按照指定格式解析一个字符串格式的时间
	// 按照东八区的时区和格式解析一个字符串格式的时间
	// 根据字符串加载时区
	loc,_ := time.LoadLocation("Asia/Shanghai")
	// 按照指定时区解析时间
	timeObj2,_ :=time.ParseInLocation("2006-01-02", "2020-05-18", loc)
	fmt.Println(timeObj2)
	fmt.Println(timeObj2.Sub(now))
	// 定时器
	/*timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}*/
	// 格式化时间 把语言中的时间对象，转换成字符串类型的时间
	// 2020-05-17
	fmt.Println(now.Format("2006-01-02"))
	// 2020/05/17 17:30:30
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2020/05/17 17:30:30 AM
	fmt.Println(now.Format("2006/01/02 03:04:05 AM"))
	// 2020/05/17 17:30:30.435
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 按照对应的格式，解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2020-05-17")
	if err != nil {
		fmt.Printf("parse time failed, err:%v", err)
		return
	}
	fmt.Println(timeObj, timeObj.Unix())
	// Sleep
	n := 5	// int
	fmt.Println("开始Sleep")
	time.Sleep(time.Duration(n) * time.Second)
	//time.Sleep(100)		// int64的纳秒
	fmt.Println("5秒过去了。。。")
	//time.Sleep(5 * time.Second)

}
