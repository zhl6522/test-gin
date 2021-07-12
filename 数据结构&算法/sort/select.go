package main

import "fmt"

func SelectSort(arr *[5]int) {
	//标准的访问方式
	//(*arr)[1] = 60 //等价于下面这个
	//arr[1] = 600

	for j := 0; j < len(arr)-1; j++ {

		//1、假设arr[0]最大值
		max := arr[j]
		maxIndex := j
		//2、遍历后面 j+1---[len(arr)-1]比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("每%d次的结果：%v\n", j+1, *arr)
	}
}

//选择排序
func main() {
	arr := [5]int{14, 60, 55, 80, 33}
	fmt.Println("原始数据：", arr)
	SelectSort(&arr)
	fmt.Println("最终结果：", arr)
}
