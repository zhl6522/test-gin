package main

import "fmt"

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		//完成第一次，给第二个元素找到合适的位置并插入
		inserVal := arr[i]
		insertIndex := i - 1 //下标
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < inserVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = inserVal
		}
		fmt.Printf("第%d次插入后 %v\n", i, *arr)
	}
}

//插入排序
func main() {
	arr := [5]int{14, 60, 55, 80, 33}
	fmt.Println("原始数据：", arr)
	InsertSort(&arr)
	fmt.Println("最终结果：", arr)
}
