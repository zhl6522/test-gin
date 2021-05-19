package main

import "fmt"

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
	fmt.Println("------------稀疏数组------------")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:\t%d\t%d\t%d\n", i, valNode.Row, valNode.Col, valNode.Val)
	}
	//写文件(写盘)参考test-gin/filedemo.go
	//读文件就不需要转结构体了，直接展示成图

	//将这个稀疏数组存盘：d:/chessmap.data

	//如何恢复原始的数组

	//1、打开这个d:/chessmap.data

	//1
}
