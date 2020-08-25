package main

import (
	"fmt"
	"time"
)

// https://leetcode-cn.com/problems/create-target-array-in-the-given-order/solution/go-shuang-100-by-ba-xiang-5/
func main() {
	st := time.Now()
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	//fmt.Println(createTargetArray([]int{2, 1, 4, 3, 0}, []int{0, 2, 1, 3, 4}))
	fmt.Println(createTargetArray([]int{1}, []int{0}))
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(index))
	sortedNum := 0 //当前已经排列总数
	for i := 0; i < len(nums); i++ {
		v := index[i]
		//fmt.Printf("&%v.%v& ", i, v)
		if i == v {
			res[i] = nums[i]
			//fmt.Printf("-1- ")
		} else {
			for n := sortedNum; n > v; n-- { //需要把数字往后移动
				res[n] = res[n-1]
				//fmt.Printf("-2- ")
			}
			res[v] = nums[i]
		}
		//fmt.Println(res[v], v)
		sortedNum++
	}
	return res
}
