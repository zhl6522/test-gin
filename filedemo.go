package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	dir := "E:/www/go_project/src/test-gin/README.md"
	fmt.Println("第一种方法，ioutil.ReadFile一次性将文件读取到位，读取不大文件更多使用的方法")
	readFile, err2 := ioutil.ReadFile(dir)
	if err2 != nil {
		fmt.Println("open file error=", err2)
	}
	//读取文件显示到终端
	//fmt.Printf("%v\n", readFile)	// []byte
	fmt.Printf("%v\n", string(readFile))
	//我们没有显式的Open文件，因此也不需要显式的Close文件
	//因为文件的Open和Close被封装到 ReadFile 函数内部了

	fmt.Println("第二种方法，读取大/较大文件更多使用的方法")
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("open file error=", err)
	}

	//当函数退出时，要及时关闭file句柄，否则会有内存泄露
	defer file.Close()

	//创建一个 *Reader，是带缓冲的
	/*
		const (
			defaultBufSize = 4096	//默认的缓冲区为4096字节
		)
	*/
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	fmt.Println("-----------写文件，window下权限设置无效-----------")
	openFile, err3 := os.OpenFile("E:/www/go_project/src/test-gin/file.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err3 != nil {
		fmt.Println("open file error=", err)
	}
	defer openFile.Close()
	str := "Hello World!\r\n"
	//写入时，使用带缓存的  *Writer
	writer := bufio.NewWriter(openFile)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用NewWriter方法时，其实内容是先写入缓存的，需要调用flush方法真正写入到文件
	writer.Flush()
}
