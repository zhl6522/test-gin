package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ValNode struct {
	Row int
	Col int
	Val int
}

func main() {
	//1、创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	//2、输出看看原始的数组
	fmt.Println("原始数据...")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3、转成稀疏数组。思想->算法
	//(1)、遍历chessMap，如果我们发现有一个元素的值不为0，创建一个结构体
	//(2)、将其放入到对应的切片即可
	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个：记录元素的二维数组的规模（行和列，默认值）
	//创建一个ValNode的值节点
	valNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode的值节点
				valNode = ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("\n------------稀疏数组------------")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:\t%d\t%d\t%d\n", i, valNode.Row, valNode.Col, valNode.Val)
	}
	//写文件(写盘)参考test-gin/filedemo.go
	//读文件就不需要转结构体了，直接展示成图
	//将这个稀疏数组存盘：d:/chessmap.data
	fmt.Println("\n<存盘>数据存入文件成功...")
	write(sparseArr)

	//如何恢复原始的数组
	//1、打开这个d:/chessmap.data
	//2、这里使用稀疏数组恢复
	read()


}

func read() {
	dir := "E:/www/go_project/src/test-gin/数据结构&算法/sparsearray/main.data"
	//fmt.Println("第一种方法，ioutil.ReadFile一次性将文件读取到位，读取不大文件更多使用的方法")
	readFile, err2 := ioutil.ReadFile(dir)
	if err2 != nil {
		fmt.Println("open file error=", err2)
	}
	var valNode []ValNode
	json.Unmarshal([]byte(readFile), &valNode)

	//先创建一个原始数组
	var chessMap2 [11][11]int
	//遍历sparseArr [遍历文件每一行]
	for i, valNode := range valNode {
		if i == 0 {
			continue
		}
		chessMap2[valNode.Row][valNode.Col] = valNode.Val
	}
	fmt.Println("\n<读盘>文件中恢复的原始数据...")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}


func write(sparseArr []ValNode) {
	openFile, err := os.OpenFile("E:/www/go_project/src/test-gin/数据结构&算法/sparsearray/main.data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error=", err)
	}
	defer openFile.Close()
	//写入时，使用带缓存的  *Writer
	writer := bufio.NewWriter(openFile)
	data, _ := json.Marshal(sparseArr)
	writer.WriteString(string(data))

	/*for _, valNode := range sparseArr {
		data := fmt.Sprintf("%d\t%d\t%d\n", valNode.Row, valNode.Col, valNode.Val)
		writer.WriteString(data)
	}*/
	//因为writer是带缓存的，因此在调用NewWriter方法时，其实内容是先写入缓存的，需要调用flush方法真正写入到文件
	writer.Flush()
}
