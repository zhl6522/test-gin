package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
// 文件三种读取方法
func readFromFile1() {
	//var fileobj *os.File	//指针的file类型
	//var err error
	//fileobj,err = os.Open("./fmt.go")
	fileobj,err := os.Open("./fmt1.go")
	//defer fileobj.Close()		// 不能写在这里的原因：os.open返回是多个值，如果打开报错了，err是非nil，fileobj是对应类型的零值（当前是指针类型的nil），nil不能调用Close(),就会panic（空指针的引用）。
	if err != nil {
		fmt.Printf("open file failed, error:%v", err)
		return
	}
	// 记得关闭文件
	defer fileobj.Close()
	// 读文件
	//var tmp = make([]byte, 128)		// 指定读的长度
	var tmp [128]byte
	for true {
		n, err := fileobj.Read(tmp[:])
		/*if err == io.EOF {
			fmt.Println("读完了")
			return
		}*/
		if err != nil {
			fmt.Printf("read file failed, error:%v", err)
			return
		}
		fmt.Printf("读了%v个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
// 利用bufio这个包读取文件
func readFromFilebyBufio() {
	fileobj,err := os.Open("./fmt.go")
	if err != nil {
		fmt.Printf("open file failed, error:%v", err)
		return
	}
	// 记得关闭文件
	defer fileobj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileobj)		// NewReader()的参数就是接口类型
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
			return
		}
		fmt.Print(line)
	}
}
func readFromFilebyIoutil() {
	ret, err := ioutil.ReadFile("./fmt.go")
	if err != nil {
		fmt.Printf("read line failed, err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

// 打开文件写内容
func writeDemo() {
	fileObj, err := os.OpenFile("./xx.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// wrie
	fileObj.Write([]byte("uu mengbi le!\n"))
	// writestring
	fileObj.WriteString("uu解释不了!")
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写得对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") //写到缓存中
	wr.Flush()		// 将缓存中的内容写入文件
}
func writeDemo3() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s,_ = reader.ReadString('\n')
	fmt.Printf("输入的内容是：%v\n", s)

}
func main() {
	//readFromFile1()
	//readFromFilebyBufio()
	//readFromFilebyIoutil()
	//writeDemo()
	//writeDemo2()
	//writeDemo3()
	useBufio()
}
