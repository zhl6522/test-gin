package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(left, right int, arr *[8000000]int) {
	l := left
	r := right
	//pivot是中轴，指点
	pivot := arr[(left+right)/2]
	//for循环的目标是将比pivot小的数放到左边，比pivot大的数放到右边
	for ; l < r; {
		//从pivot的左边找到大于等于pivot的值
		for ; arr[l] < pivot; {
			l++
		}
		//从pivot的右边找到小于等于pivot的值
		for ; arr[r] > pivot; {
			r--
		}
		//l>=r表明本次分解任务完成
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		//优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//fmt.Printf("l=%d r=%d\n", l, r)
	//如果l==r，在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

//快速排序	800W数据1-2s完成
func main() {
	var arr [8000000]int
	for i := 0; i < 8000000; i++ {
		arr[i] = rand.Intn(90000000)
	}
	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr) //特别注意：代码里面不要有输出，不然输出时间也会被计算在内，这是不合理的
	end := time.Now().Unix()
	fmt.Printf("快速排序耗时%d秒", end-start)
	fmt.Println(arr[7999999], arr[7999998], arr[7999997])

	//arr := [6]int{-9, 78, 0, 23, -567, 70}
	//fmt.Println("原始数据：", arr)
	//QuickSort(0, len(arr)-1, &arr)
	//fmt.Println("最终结果：", arr)
}
