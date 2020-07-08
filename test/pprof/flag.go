package main

import (
	"flag"
	"fmt"
)

// flag 获取命令行参数
func main() {
	// 创建一个标志位参数
	name := flag.String("name", "司仪", "请输入名字")
	age := flag.Int("age", 26, "请输入真实年龄")
	// Bool false; Duration time.Second

	// 使用flag
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)

	/*var name string
	flag.StringVar(&name, "name", "司仪", "请输入名字")
	// 使用flag
	flag.Parse()
	fmt.Println(name)*/

	fmt.Println(flag.Args())  ////返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数

}
