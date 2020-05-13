package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	fileobj,err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %v failed, error:%v", srcName, err)
		return
	}
	fmt.Println(fileobj)
	// 记得关闭文件
	defer fileobj.Close()
	fileObj, err := os.OpenFile(dstName,os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	return io.Copy(fileObj, fileobj)
}

func main() {
	_,err := CopyFile("./dst.txt", "./src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}
