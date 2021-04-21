package main

import (
	"fmt"
)

//冒泡排序
func main() {
	//append操作可能会导致原本使用同一个底层数组的两个Slice变量变为使用不同的底层数组。
	a := [3]int{22, 59, 43}
	slice := a[:]
	fmt.Printf("%p\n",slice)
	slice = append(slice, 1,2)
	fmt.Printf("%p\n",slice)


	arr := [6]int{22, 59, 43, 77, 16, 7}

	BubbleAsort(arr[:])		//切片解决冒泡排序	常用切片

	//ArrayAsort(&arr) //纯数组解决冒泡排序
	//fmt.Printf("排序后的数组：%v\n", arr)
	//Array2Asort(arr) //纯数组解决冒泡排序

}

func BubbleAsort(values []int) {
	fmt.Printf("排序前的数组：%v\n", values)
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
		fmt.Println(values, "-------")
	}
	fmt.Printf("排序前的数组：%v\n", values)
}

func ArrayAsort(arr *[6]int) {
	fmt.Printf("排序前的数组：%v\n", *arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[j],(*arr)[i] = (*arr)[i],(*arr)[j]
			}
		}
	}
}

func Array2Asort(arr [6]int) {
	var k int
	var str int
	fmt.Printf("排序前的数组：%v\n", arr)
	for n := 0; n < len(arr); n++ {
		num := arr[n]
		k = n
		for i := n + 1; i < len(arr); i++ {
			if num > arr[i] {
				if str != 0 && num <= str {
					continue
				}
				k = i
				num = arr[i]
			}
		}
		str = arr[k]
		fmt.Printf("数组中最小值的下标：%v 替换值的下标：%v\n", k, n)
		arr[k], arr[n] = arr[n], num
	}
	fmt.Println(arr)
}
