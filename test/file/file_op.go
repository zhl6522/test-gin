package main

import (
	"fmt"
	"io"
	"os"
)

func f() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./seek.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//defer fileObj.Close()
	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./seek.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, err:%v", err)
		return
	}
	defer tmpFile.Close()
	// 读原文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	// 写入临时文件
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	// 紧接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:]);
		if err == io.EOF {
			tmpFile.Write(x[:n])
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	// 原文件后续的也写入临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./seek.tmp", "./seek.txt")

	//fmt.Println(string(ret[:n]))
	//fileObj.Seek(3, 0)	// 光标移动到b（是3的原因：windows换行是：\n\r）


}

func main() {
	f()
}