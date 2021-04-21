package main

import "fmt"

//二分查找
func Binaryfind(arr *[6]int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到。。。")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		Binaryfind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		Binaryfind(arr, middle+1, rightIndex,findVal)
	} else {
		fmt.Printf("找到了，下标为%v\n", middle)
	}
}

func main() {
	var num int
	arr := [6]int{5, 16, 27, 38, 49, 66}
	fmt.Printf("当前的数组列表：%v\n", arr)
	fmt.Println("请选择要查找的数字：")
	fmt.Scanln(&num)
	Binaryfind(&arr, 0, len(arr)-1, num)
}
